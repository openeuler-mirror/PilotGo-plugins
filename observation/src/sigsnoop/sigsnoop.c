// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "sigsnoop.h"
#include "sigsnoop.skel.h"
#include "compat.h"

#include <libgen.h>

static volatile sig_atomic_t exiting;

static pid_t target_pid = 0;
static int target_signal = 0;
static bool failed_only = false;
static bool kill_only = false;
static bool signal_name = false;
static bool verbose = false;

int main(int argc, char *argv[])
{
	static const struct argp argp = {
		.options = opts,
		.parser = parse_arg,
		.doc = argp_program_doc,
	};
	int err;

	alias_parse(argv[0]);
	err = argp_parse(&argp, argc, argv, 0, NULL, NULL);
	if (err)
		return err;
	
	return err != 0;
}

