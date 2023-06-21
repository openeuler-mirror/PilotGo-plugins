// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "biolatency.h"
#include "biolatency.skel.h"
#include "trace_helpers.h"
#include "blk_types.h"
#include <sys/resource.h>

static struct env {
	char	*disk;
	time_t	interval;
	int	times;
	bool	timestamp;
	bool	queued;
	bool	per_disk;
	bool	per_flag;
	bool	milliseconds;
	bool	verbose;
	char	*cgroupspath;
	bool	cg;
} env = {
	.interval = 99999999,
	.times = 99999999,
};

static volatile sig_atomic_t exiting;

const char *argp_program_version = "biolatency 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Summarize block device I/O latency as a histogram.\n"
"\n"
"USAGE: biolatency [--help] [-T] [-m] [-Q] [-D] [-F] [-d DISK] [-c CG] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    biolatency              # summarize block I/O latency as a histogram\n"
"    biolatency 1 10         # print 1 second summaries, 10 times\n"
"    biolatency -mT 1        # 1s summaries, milliseconds, and timestamps\n"
"    biolatency -Q           # include OS queued time in I/O time\n"
"    biolatency -D           # show each disk device separately\n"
"    biolatency -F           # show I/O flags separately\n"
"    biolatency -d sdc       # Trace sdc only\n"
"    biolatency -c CG        # Trace process under cgroupsPath CG\n";

static const struct argp_option opts[] = {
	{ "timestamp", 'T', NULL, 0, "Include timestamp on output" },
	{ "milliseonds", 'm', NULL, 0, "Millisecond histogram" },
	{ "queued", 'Q', NULL, 0, "Include OS queued time in I/O time" },
	{ "disk", 'D', NULL, 0, "Print a histogram per disk device" },
	{ "flag", 'F', NULL, 0, "Print a histogram per set of I/O flags" },
	{ "disk", 'd', "DISK", 0, "Trace this disk only" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ "cgroup", 'c', "/sys/fs/cgroup/unified", 0, "Trace process in cgroup path" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{},
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
	case 'T':
		env.timestamp = true;
		break;
	case 'm':
		env.milliseconds = true;
		break;
	case 'Q':
		env.queued = true;
		break;
	case 'D':
		env.per_disk = true;
		break;
	case 'F':
		env.per_flag = true;
		break;
	case 'c':
		env.cgroupspath = arg;
		env.cg = true;
		break;
	case 'd':
		env.disk = arg;
		if (strlen(arg) + 1 > DISK_NAME_LEN) {
			warning("Invalid disk name: too long\n");
			argp_usage(state);
		}
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
