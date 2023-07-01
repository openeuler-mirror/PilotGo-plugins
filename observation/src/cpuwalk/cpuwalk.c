// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "cpuwalk.h"
#include "cpuwalk.skel.h"
#include "trace_helpers.h"
#include "btf_helpers.h"
#include <sys/syscall.h>
#include <linux/perf_event.h>

static volatile sig_atomic_t exiting;

static bool verbose = false;
static int frequency = 99;

const char *argp_program_version = "cpuwalk 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Sample which CPUs are executing processes\n"
"\n"
"USAGE: cpuwalk [-v] [-f FREQUENCY]\n"
"\n"
"Example:\n"
"    cpuwalk               # sampling cpu\n"
"    cpuwalk -f 199        # sampling at 199HZ\n";

static const struct argp_option opts[] = {
	{ "frequency", 'f', "FREQUENCY", 0, "Sample with a certain frequency" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};


static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	switch (key) {
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		verbose = true;
		break;
	case 'f':
		frequency = argp_parse_long(key, arg, state);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}
	return 0;
}

static int nr_cpus;

static int open_and_attach_perf_event(struct bpf_program *prog,
				      struct bpf_link *links[])
{
	for (int i = 0; i < nr_cpus; i++) {
		struct perf_event_attr attr = {
			.type = PERF_TYPE_SOFTWARE,
			.freq = 1,
			.sample_freq = frequency,
			.config = PERF_COUNT_SW_CPU_CLOCK,
		};

		int fd = syscall(__NR_perf_event_open, &attr, -1, i, -1, 0);
		if (fd < 0) {
			/* Ignore CPU that is offline */
			if (errno == ENODEV)
				continue;

			warning("Failed to init perf sampling: %s\n", strerror(errno));
			return -1;
		}

		links[i] = bpf_program__attach_perf_event(prog, fd);
		if (!links[i]) {
			warning("Failed to attach perf event on CPU #%d!\n", i);
			close(fd);
			return -1;
		}
	}

	return 0;
}

static int libbpf_print_fn(enum libbpf_print_level level, const char *format,
			   va_list args)
{
	if (level == LIBBPF_DEBUG && !verbose)
		return 0;
	return vfprintf(stderr, format, args);
}

static void sig_handler(int sig)
{
	exiting = 1;
}

static struct hist zero;

static void print_hist(struct cpuwalk_bpf__bss *bss)
{
	struct hist hist = bss->hist;

	printf("\n");

	bss->hist = zero;
	print_linear_hist(hist.slots, MAX_CPU_NR, 0, 1, "cpuwalk");
}
