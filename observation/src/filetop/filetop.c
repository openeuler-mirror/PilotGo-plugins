// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "filetop.h"
#include "filetop.skel.h"
#include "btf_helpers.h"
#include "trace_helpers.h"

#define OUTPUT_ROWS_LIMIT	10240

enum SORT {
	ALL,
	READS,
	WRITES,
	RBYTES,
	WBYTES,
};

static volatile sig_atomic_t exiting;
static volatile bool verbose = false;
static volatile int sort_by = ALL;

struct argument {
	pid_t target_pid;
	bool clear_screen;
	bool regular_file_only;
	int output_rows;
	int interval;
	int count;
};

const char *argp_program_version = "filetop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace file reads/writes by process.\n"
"\n"
"USAGE: filetop [-h] [-p PID] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    filetop            # file I/O top, refresh every 1s\n"
"    filetop -p 1216    # only trace PID 1216\n"
"    filetop 5 10       # 5s summaries, 10 times\n";

static const struct argp_option opts[] = {
	{ "pid", 'p', "PID", 0, "Process ID to trace" },
	{ "noclear", 'C', NULL, 0, "Don't clear the screen" },
	{ "all", 'a', NULL, 0, "Include special files" },
	{ "sort", 's', "SORT", 0, "Sort columns, default all [all, reads, writes, rbytes, wbytes]" },
	{ "rows", 'r', "ROWS", 0, "Maximum rows to print, default 20" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	struct argument *argument = state->input;
	static int pos_args;
	int max_rows = OUTPUT_ROWS_LIMIT;

	switch (key) {
	case 'p':
		argument->target_pid = argp_parse_pid(key, arg, state);
		break;
	case 'C':
		argument->clear_screen = false;
		break;
	case 'a':
		argument->regular_file_only = false;
		break;
	case 's':
		if (!strcmp(arg, "all"))
			sort_by = ALL;
		else if (!strcmp(arg, "reads"))
			sort_by = READS;
		else if (!strcmp(arg, "writes"))
			sort_by = WRITES;
		else if (!strcmp(arg, "rbytes"))
			sort_by = RBYTES;
		else if (!strcmp(arg, "wbytes"))
			sort_by = WBYTES;
		else {
			warning("Invalid sort method: %s\n", arg);
			argp_usage(state);
		}
		break;
	case 'r':
		errno = 0;
		argument->output_rows = strtol(arg, NULL, 10);
		if (errno || argument->output_rows <= 0) {
			warning("Invalud rows: %s\n", arg);
			argp_usage(state);
		}
		argument->output_rows = min(max_rows, argument->output_rows);
		break;
	case 'v':
		verbose = true;
		break;
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
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
			argument->count = strtol(arg, NULL, 10);
			if (errno || argument->count <= 0) {
				warning("Invalid count\n");
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
	if (level == LIBBPF_DEBUG && !verbose)
		return 0;
	return vfprintf(stderr, format, args);
}

static void sig_handler(int sig)
{
	exiting = 1;
}

static int sort_column(const void *obj1, const void *obj2)
{
	struct file_stat *s1 = (struct file_stat *)obj1;
	struct file_stat *s2 = (struct file_stat *)obj2;

	if (sort_by == READS) {
		return s2->reads - s1->reads;
	} else if (sort_by == WRITES) {
		return s2->writes - s1->writes;
	} else if (sort_by == RBYTES) {
		return s2->read_bytes - s1->read_bytes;
	} else if (sort_by == WBYTES) {
		return s2->write_bytes - s1->write_bytes;
	} else {
		return (s2->reads + s2->writes + s2->read_bytes + s2->write_bytes) -
			(s1->reads + s1->writes + s1->read_bytes + s1->write_bytes);
	}
}
