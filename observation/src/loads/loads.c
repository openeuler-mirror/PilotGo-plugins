// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "loads.skel.h"
#include "btf_helpers.h"
#include "trace_helpers.h"
#include <linux/perf_event.h>
#include <sys/syscall.h>

static struct env
{
	int interval;
	int times;
	bool verbose;
	bool timestamp;
} env = {
	.times = 99999999,
	.interval = 1,
};

#define MAX_NR_CPUS 1024

static volatile sig_atomic_t exiting;

const char *argp_program_version = "loads 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
	"Print load averages\n"
	"\n"
	"USAGE: loads [-i INTERVAL] [-t times]\n"
	"\n"
	"EXAMPLE:\n"
	"    loads                 # print load average every 1 seconds\n"
	"    loads -i 10           # print load average every 10 seconds\n"
	"    loads -t 5            # print load average 5 times\n";

static const struct argp_option opts[] = {
	{"interval", 'i', "INTERVAL", 0, "Output interval, in seconds (Default 1)"},
	{"times", 't', "TIMES", 0, "The number of outputs"},
	{"verbose", 'v', NULL, 0, "Verbose debug output"},
	{"timestamp", 'T', NULL, 0, "Include timestamp on output"},
	{NULL, 'h', NULL, OPTION_HIDDEN, "SHow the full help"},
	{}};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	switch (key)
	{
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		env.verbose = true;
		break;
	case 'T':
		env.timestamp = true;
		break;
	case 't':
		env.times = argp_parse_long(key, arg, state);
		break;
	case 'i':
		env.interval = argp_parse_long(key, arg, state);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}
	return 0;
}

static int nr_cpus;

static int open_and_attach_perf_event(struct bpf_program *prog, struct bpf_link *link[])
{
	for (int i = 0; i < nr_cpus; i++)
	{
		struct perf_event_attr attr = {
			.type = PERF_TYPE_SOFTWARE,
			.config = PERF_COUNT_SW_CPU_CLOCK,
			.sample_period = 1 / env.interval,
			.freq = 1,
		};

		int fd = syscall(SYS_perf_event_open, &attr, -1, i, -1, 0);
		if (fd < 0)
		{
			/* Ignore CPU that is offline */
			if (errno == ENODEV)
				continue;

			warning("Failed to init perf sampling: %s\n", strerror(errno));
			return -1;
		}

		link[i] = bpf_program__attach_perf_event(prog, fd);
		if (!link[i])
		{
			warning("Failed to attach perf event on CPU#0\n");
			close(fd);
			return 1;
		}
	}

	return 0;
}

static int libbpf_print_fn(enum libbpf_print_level level, const char *format,
						   va_list args)
{
	if (level == LIBBPF_DEBUG && !env.verbose)
		return 0;
	return vfprintf(stderr, format, args);
}

static void sig_handler(int sig)
{
	exiting = 1;
}

static void print_loads(struct loads_bpf__bss *bss)
{
	__u64 load1 = __atomic_load_n(&bss->loads[0], __ATOMIC_RELAXED);
	__u64 load5 = __atomic_load_n(&bss->loads[1], __ATOMIC_RELAXED);
	__u64 load15 = __atomic_load_n(&bss->loads[2], __ATOMIC_RELAXED);

	if (env.timestamp)
	{
		char ts[32];

		strftime_now(ts, sizeof(ts), "%H:%M:%S");
		printf("%s ", ts);
	}

	printf("load averages: %lld.%03lld %lld.%03lld %lld.%03lld\n",
		   load1 >> 11, ((load1 & ((1 << 11) - 1)) * 1000) >> 11,
		   load5 >> 11, ((load5 & ((1 << 11) - 1)) * 1000) >> 11,
		   load15 >> 11, ((load15 & ((1 << 11) - 1)) * 1000) >> 11);
}