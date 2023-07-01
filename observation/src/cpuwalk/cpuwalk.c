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
