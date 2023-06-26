// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
// Copyright @ 2023 - Kylin
// Author: Jackie Liu <liuyun01@kylinos.cn>

#include "commons.h"
#include "compat.h"
#include "filesnoop.h"
#include "filesnoop.skel.h"
#include "trace_helpers.h"

static volatile sig_atomic_t exiting;

static struct env {
	bool verbose;
	bool timestamp;
	const char *filename;
	bool filter_filename;
	enum file_op target_op;
} env;

const char *argp_program_version = "filesnoop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Tracking the operational of a specific file.\n"
"\n"
"USAGE: filesnoop [-v] [-T] [-f filename] [-o OPEN]\n"
"\n"
"EXAMPLE:\n"
"    filesnoop -o OPEN        # trace open/openat/openat2 syscall\n"
"                             # (open,write,read,stat,close)\n";
