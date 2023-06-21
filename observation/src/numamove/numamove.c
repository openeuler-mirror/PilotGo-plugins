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

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	switch (key)
	{
	case 'v':
		env.verbose = true;
		break;
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}

	return 0;
}

static int libbpf_print_fn(enum libbpf_print_level level, const char *format,
						   va_list args)
{
	if (level == LIBBPF_DEBUG && !env.verbose)
		return 0;

	return vfprintf(stderr, format, args);
}

static void sig_handler(int sig)
{
	exiting = 1;
}