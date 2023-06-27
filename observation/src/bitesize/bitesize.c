// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "bitesize.h"
#include "bitesize.skel.h"
#include "trace_helpers.h"

struct argument {
	char *disk;
	char *comm;
	int comm_len;
	time_t interval;
	bool timestamp;
	int times;
};

static volatile bool verbose = false;
static volatile sig_atomic_t exiting;

const char *argp_program_version = "bitesize 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Summarize block device I/O size as a histogram.\n"
"\n"
"USAGE: bitesize [--help] [-T] [-c COMM] [-d DISK] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    bitesize              # summarize block I/O latency as a histogram\n"
"    bitesize 1 10         # print 1 second summaries, 10 times\n"
"    bitesize -T 1         # 1s summaries with timestamps\n"
"    bitesize -c fio       # trace fio only\n";

static const struct argp_option opts[] = {
	{ "timestamp", 'T', NULL, 0, "Include timestamp on output" },
	{ "comm", 'c', "COMM", 0, "Trace this comm only" },
	{ "disk", 'd', "DISK", 0, "Trace this disk only" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

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
