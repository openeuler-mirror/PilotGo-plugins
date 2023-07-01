// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "dcsnoop.h"
#include "dcsnoop.skel.h"
#include "trace_helpers.h"
#include "compat.h"

static volatile sig_atomic_t exiting;
static __u64 time_end;

static struct env {
	bool trace_all;
	bool verbose;
	bool timestamp;
	bool duration;
	pid_t pid;
	pid_t tid;
} env;

const char *argp_program_version = "dcsnoop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace directory entry cache (dcache) lookups.\n"
"\n"
"USAGE: dcsnoop [-a] [-T] [-d DURATION] [-v] [-t TID] [-p PID]\n"
"\n"
"Examples: \n"
"    dcsnoop               # trace failed dcache lookups\n"
"    dcsnoop -a            # trace all dcache lookups\n"
"    dcsnoop -T            # include timestamp on output\n"
"    dcsnoop -d 10         # 10s to trace\n"
"    dcsnoop -p 188        # trace pid 188 only\n"
"    dcsnoop -t 188        # trace tid 188 only\n";

static const struct argp_option opts[] = {
	{ "trace-all", 'a', NULL, 0, "Trace all lookups (default is fails only)" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ "timestamp", 'T', NULL, 0, "Include timestamp on output" },
	{ "duration", 'd', "DURATION", 0, "Duration to trace" },
	{ "pid", 'p', "PID", 0, "Trace process ID only" },
	{ "tid", 't', "TID", 0, "Trace thread ID only" },
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
		env.verbose = true;
		break;
	case 'a':
		env.trace_all = true;
		break;
	case 'T':
		env.timestamp = true;
		break;
	case 'd':
		env.duration = argp_parse_long(key, arg, state);
		break;
	case 'p':
		env.pid = argp_parse_pid(key, arg, state);
		break;
	case 't':
		env.tid = argp_parse_pid(key, arg, state);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}
	return 0;
}
