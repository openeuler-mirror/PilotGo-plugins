// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
// Copyright @ 2023 - Kylin
// Author: Jackie Liu <liuyun01@kylinos.cn>
//
// Idea by Brendan Gregg.
// Base on bpflist-bpfcc(8) - Copyright 2017, Sasha Goldshtein
#include "commons.h"
#include <dirent.h>
#include <regex.h>

static int verbose = 0;

const char *argp_program_version = "bpflist 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Display processes currently using BPF programs and maps, pinned BPF programs"
" and maps, and enabled probes.\n"
"\n"
"USAGE: bpflist [-v]\n";

static const struct argp_option opts[] = {
	{ "verbose", 'v', NULL, 0, "also count kprobes/uprobes" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};
