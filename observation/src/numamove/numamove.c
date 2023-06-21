// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "numamove.skel.h"
#include "trace_helpers.h"
#include <numa.h>

static struct env
{
	bool verbose;
} env = {
	.verbose = false,
};

static volatile sig_atomic_t exiting;

const char *argp_program_version = "numamove 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
	"Show page migrations of type NUMA misplaced per second.\n"
	"\n"
	"USAGE: numamove [--help]\n"
	"\n"
	"EXAMPLES:\n"
	"    numamove              # Show page migrations' count and latency";

static const struct argp_option opts[] = {
	{"verbose", 'v', NULL, 0, "Verbose debug output"},
	{NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help"},
	{}};