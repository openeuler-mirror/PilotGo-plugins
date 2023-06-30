// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "cpufreq.h"
#include "cpufreq.skel.h"
#include "trace_helpers.h"
#include <linux/perf_event.h>
#include <sys/syscall.h>

int main(int argc, char *argv[])
{
	static const struct argp argp = {
		.parser = parse_arg,
		.options = opts,
		.doc = argp_program_doc,
	};
	struct bpf_link *links[MAX_CPU_NR] = {};
	struct cpufreq_bpf *obj;
	int err, cgfd = -1;

	err = argp_parse(&argp, argc, argv, 0, NULL, NULL);
	if (err)
		return err;

	if (!bpf_is_root())
		return 1;

	libbpf_set_print(libbpf_print_fn);

	nr_cpus = libbpf_num_possible_cpus();
	if (nr_cpus < 0) {
		warning("Failed to get # of possible cpus: '%s'!\n",
			strerror(-nr_cpus));
		return 1;
	}

	if (nr_cpus > MAX_CPU_NR) {
		warning("the number of cpu cores is too big, please "
			"increase MAX_CPU_NR's value and recompile");
		return 1;
	}

	obj = cpufreq_bpf__open();
	if (!obj) {
		warning("Failed to open BPF object\n");
		return 1;
	}

	if (probe_tp_btf("cpu_frequency"))
		bpf_program__set_autoload(obj->progs.cpu_frequency_raw, false);
	else
		bpf_program__set_autoload(obj->progs.cpu_frequency, false);

	err = cpufreq_bpf__load(obj);
	if (err) {
		warning("Failed to load BPF object\n");
		goto cleanup;
	}

	if (!obj->bss) {
		warning("Memory-mapping BPF maps is supported starting from Linux 5.7, please upgrade.\n");
		goto cleanup;
	}

	err = init_freqs_mhz(obj->bss->freqs_mhz, nr_cpus);
	if (err) {
		warning("Failed to init freqs\n");
		goto cleanup;
	}

	obj->bss->filter_memcg = env.cg;

	/* update cgroup path fd to map */
	if (env.cg) {
		int idx = 0;
		int cg_map_fd = bpf_map__fd(obj->maps.cgroup_map);

		cgfd = open(env.cgroupspath, O_RDONLY);
		if (cgfd < 0) {
			warning("Failed opening Cgroup path: %s", env.cgroupspath);
			goto cleanup;
		}
		if (bpf_map_update_elem(cg_map_fd, &idx, &cgfd, BPF_ANY)) {
			warning("Failed adding target cgroup to map");
			goto cleanup;
		}
	}

	err = open_and_attach_perf_event(env.freq, obj->progs.do_sample, links);
	if (err)
		goto cleanup;

	err = cpufreq_bpf__attach(obj);
	if (err) {
		warning("Failed to attach BPF programs\n");
		goto cleanup;
	}

	printf("Sampling CPU freq system-wide & by process. Ctrl-C to end.\n");

	signal(SIGINT, sig_handler);

	/*
	 * we'll get sleep interrupted when someone process Ctrl-C (which will
	 * be "handled" with noop by sig_handler).
	 */
	sleep(env.duration);
	printf("\n");

	print_linear_hists(obj->maps.hists, obj->bss);

cleanup:
	for (int i = 0; i < nr_cpus; i++)
		bpf_link__destroy(links[i]);
	cpufreq_bpf__destroy(obj);
	if (cgfd > 0)
		close(cgfd);

	return err != 0;
}
