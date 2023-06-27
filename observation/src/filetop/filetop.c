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
