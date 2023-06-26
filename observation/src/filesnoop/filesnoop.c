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

static const struct argp_option opts[] = {
	{ "version", 'v', NULL, 0, "Verbose debug output" },
	{ "timestamp", 'T', NULL, 0, "Include timestamp on output" },
	{ "filename", 'f', "FILENAME", 0, "Trace FILENAME only" },
	{ "operation", 'o', "OPERATION", 0, "Trace OPERATION only" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

const char *op2string[] = {
	[F_ALL] = "NONE",
	[F_OPEN] = "OPEN",
	[F_OPENAT] = "OPENAT",
	[F_OPENAT2] = "OPENAT2",
	[F_WRITE] = "WRITE",
	[F_WRITEV] = "WRITEV",
	[F_READ] = "READ",
	[F_READV] = "READV",
	[F_STATX] = "STATX",
	[F_FSTATFS] = "FSTATFS",
	[F_NEWFSTAT] = "NEWFSTAT",
	[F_RENAMEAT] = "RENAMEAT",
	[F_RENAMEAT2] = "RENAMEAT2",
	[F_UNLINKAT] = "UNLINKAT",
	[F_CLOSE] = "CLOSE",
	[F_UTIMENSAT] = "UTIMENSAT",
};
