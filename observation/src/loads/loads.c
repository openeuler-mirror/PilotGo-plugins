// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "loads.skel.h"
#include "btf_helpers.h"
#include "trace_helpers.h"
#include <linux/perf_event.h>
#include <sys/syscall.h>

static struct env
{
	int interval;
	int times;
	bool verbose;
	bool timestamp;
} env = {
	.times = 99999999,
	.interval = 1,
};

#define MAX_NR_CPUS 1024

static volatile sig_atomic_t exiting;

const char *argp_program_version = "loads 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
	"Print load averages\n"
	"\n"
	"USAGE: loads [-i INTERVAL] [-t times]\n"
	"\n"
	"EXAMPLE:\n"
	"    loads                 # print load average every 1 seconds\n"
	"    loads -i 10           # print load average every 10 seconds\n"
	"    loads -t 5            # print load average 5 times\n";

static const struct argp_option opts[] = {
	{"interval", 'i', "INTERVAL", 0, "Output interval, in seconds (Default 1)"},
	{"times", 't', "TIMES", 0, "The number of outputs"},
	{"verbose", 'v', NULL, 0, "Verbose debug output"},
	{"timestamp", 'T', NULL, 0, "Include timestamp on output"},
	{NULL, 'h', NULL, OPTION_HIDDEN, "SHow the full help"},
	{}};