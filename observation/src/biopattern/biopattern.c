// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "biopattern.h"
#include "biopattern.skel.h"
#include "btf_helpers.h"
#include "trace_helpers.h"

static struct env {
	char *disk;
	time_t interval;
	bool timestamp;
	bool verbose;
	int times;
} env = {
	.interval = 99999999,
	.times = 99999999,
};

static volatile sig_atomic_t exiting;

const char *argp_program_version = "biopattern 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Show block device I/O pattern.\n"
"\n"
"USAGE: biopattern [--help] [-T] [-d DISK] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    biopattern              # show block I/O pattern\n"
"    biopattern 1 10         # print 1 second summaries, 10 times\n"
"    biopattern -T 1         # 1s summaries with timestamps\n"
"    biopattern -d sdc       # trace sdc only\n";

static const struct argp_option opts[] = {
	{ "timestamp", 'T', NULL, 0, "Include timestamp on output" },
	{ "disk", 'd', "DISK", 0, "Trace this disk only" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	static int pos_args;

	switch (key) {
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		env.verbose = true;
		break;
	case 'd':
		env.disk = arg;
		if (strlen(arg) + 1 > DISK_NAME_LEN) {
			warning("Invalid disk name: too long\n");
			argp_usage(state);
		}
		break;
	case 'T':
		env.timestamp = true;
		break;
	case ARGP_KEY_ARG:
		errno = 0;
		if (pos_args == 0) {
			env.interval = strtol(arg, NULL, 10);
			if (errno || env.interval <= 0) {
				warning("Invalid interval\n");
				argp_usage(state);
			}
		} else if (pos_args == 1) {
			env.times = strtol(arg, NULL, 10);
			if (errno || env.times <= 0) {
				warning("Invalid times\n");
				argp_usage(state);
			}
		} else {
			warning("Unrecognized positional argument: %s\n", arg);
			argp_usage(state);
		}
		pos_args++;
		break;
	default:
		return ARGP_ERR_UNKNOWN;
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

static int print_map(struct bpf_map *counters, struct partitions *partitions)
{
	__u32 total, lookup_key = -1, next_key;
	int err, fd = bpf_map__fd(counters);
	const struct partition *partition;
	struct counter counter;

	while (!bpf_map_get_next_key(fd, &lookup_key, &next_key)) {
		err = bpf_map_lookup_elem(fd, &next_key, &counter);
		if (err < 0) {
			warning("Failed to lookup counters: %d\n", err);
			return -1;
		}

		lookup_key = next_key;
		total = counter.sequential + counter.random;
		if (!total)
			continue;
		if (env.timestamp) {
			char ts[32];

			strftime_now(ts, sizeof(ts), "%H:%M:%S");
			printf("%-9s ", ts);
		}
		partition = partitions__get_by_dev(partitions, next_key);
		printf("%-7s %5ld %5ld %8d %10lld\n",
		       partition ? partition->name : "Unknown",
		       counter.random * 100L / total,
		       counter.sequential * 100L / total, total,
		       counter.bytes / 1024);
	}

	lookup_key = -1;
	while (!bpf_map_get_next_key(fd, &lookup_key, &next_key)) {
		err = bpf_map_delete_elem(fd, &next_key);
		if (err < 0) {
			warning("Failed to cleanup counters: %d\n", err);
			return -1;
		}
		lookup_key = next_key;
	}

	return 0;
}