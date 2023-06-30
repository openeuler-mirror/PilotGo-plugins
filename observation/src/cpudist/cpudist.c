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

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	switch (key) {
	case 'v':
		verbose++;
		break;
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}

	return 0;
}

#define MAX_PATH_LEN	256

static char *comm_for_pid(const char *pid)
{
	char comm_path[MAX_PATH_LEN];
	snprintf(comm_path, sizeof(comm_path), "/proc/%s/comm", pid);

	FILE *file = fopen(comm_path, "r");
	if (!file)
		return "[unknown]";

	char *buffer = NULL;
	size_t length = 0;
	ssize_t read;

	read = getline(&buffer, &length, file);
	fclose(file);

	if (read == -1) {
		if (buffer)
			free(buffer);
		return "[unknown]";
	}

	if (buffer[read - 1] == '\n')
		buffer[read - 1] = 0;

	return buffer;
}
