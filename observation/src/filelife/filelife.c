// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "filelife.h"
#include "filelife.skel.h"
#include "btf_helpers.h"
#include "trace_helpers.h"

static volatile sig_atomic_t exiting;
static volatile bool verbose = false;
static volatile pid_t pid = 0;

const char *argp_program_version = "filelife 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace the lifespan of short-lived files.\n"
"\n"
"USAGE: filelife  [--help] [-p PID]\n"
"\n"
"EXAMPLES:\n"
"    filelife         # trace all events\n"
"    filelife -p 123  # trace pid 123\n";

static const struct argp_option opts[] = {
	{ "pid", 'p', "PID", 0, "Process PID to trace" },
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
	case 'p':
		pid = argp_parse_pid(key, arg, state);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
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

void handle_event(void *ctx, int cpu, void *data, __u32 data_sz)
{
	const struct event *e = data;
	char ts[32];

	strftime_now(ts, sizeof(ts), "%H:%M:%S");
	printf("%-8s %-6d %-16s %-7.2f %s\n",
		ts, e->tgid, e->task, (double)e->delta_ns / 1e9,
		e->file);
}