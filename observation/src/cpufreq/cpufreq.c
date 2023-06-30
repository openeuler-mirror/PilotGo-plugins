// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "cachestat.skel.h"
#include "trace_helpers.h"

struct argument {
	time_t	interval;
	int	times;
	bool	timestamp;
};

static volatile sig_atomic_t exiting;
static volatile bool verbose = false;

const char *argp_program_version = "cachestat 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Count cache kernel function calls.\n"
"\n"
"USAGE: cachestat [--help] [-T] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    cachestat          # shows hits and misses to the file system page cache\n"
"    cachestat -T       # include timestamps\n"
"    cachestat 1 10     # print 1 second summaries, 10 times\n";

static const struct argp_option opts[] = {
	{ "timestamp", 'T', NULL, 0, "Include timestamp" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	static int pos_args;
	struct argument *argument = state->input;

	switch (key) {
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		verbose = true;
		break;
	case 'T':
		argument->timestamp = true;
		break;
	case ARGP_KEY_ARG:
		errno = 0;
		if (pos_args == 0) {
			argument->interval = strtol(arg, NULL, 10);
			if (errno || argument->interval <= 0) {
				warning("Invalid interval\n");
				argp_usage(state);
			}
		} else if (pos_args == 1) {
			argument->times = strtol(arg, NULL, 10);
			if (errno || argument->times <= 0) {
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