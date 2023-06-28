// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)

#include "commons.h"
#include "biotop.h"
#include "biotop.skel.h"
#include "trace_helpers.h"
#include "compat.h"

#define OUTPUT_ROWS_LIMIT	10240

enum SORT {
	ALL,
	IO,
	BYTES,
	TIME,
};

struct disk {
	int major;
	int minor;
	char name[256];
};

struct vector {
	size_t nr;
	size_t capacity;
	void **elems;
};

int grow_vector(struct vector *vector)
{
	if (vector->nr >= vector->capacity) {
		void **reallocated;

		if (!vector->capacity)
			vector->capacity = 1;
		else
			vector->capacity *= 2;

		reallocated = libbpf_reallocarray(vector->elems, vector->capacity, sizeof(*vector->elems));
		if (!reallocated)
			return -1;

		vector->elems = reallocated;
	}

	return 0;
}

void free_vector(struct vector vector)
{
	for (size_t i = 0; i < vector.nr; i++) {
		if (vector.elems[i] != NULL)
			free(vector.elems[i]);
	}

	free(vector.elems);
}

struct vector disks = {};

static volatile sig_atomic_t exiting;

static struct env {
	bool	clear_screen;
	int	output_rows;
	int	sort_by;
	int	interval;
	int	count;
	bool	verbose;
} env = {
	.clear_screen	= true,
	.output_rows	= 20,
	.sort_by	= ALL,
	.interval	= 1,
	.count		= 99999999,
	.verbose	= false,
};

const char *argp_program_version = "biotop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace file reads/writes by process.\n"
"\n"
"USAGE: biotop [-h] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    biotop            # file I/O top, refresh every 1s\n"
"    biotop 5 10       # 5s summaries, 10 times\n";

static const struct argp_option opts[] = {
	{ "noclear", 'c', NULL, 0, "Don't clear the screen" },
	{ "sort", 's', "SORT", 0, "Sort columns, default all [all, io, bytes, time]" },
	{ "rows", 'r', "ROWS", 0, "Maximum rows to print, default 20" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	static int pos_args;

	switch (key) {
	case 'c':
		env.clear_screen = false;
		break;
	case 's':
		if (!strcmp(arg, "all")) {
			env.sort_by = ALL;
		} else if (!strcmp(arg, "io")) {
			env.sort_by = IO;
		} else if (!strcmp(arg, "bytes")) {
			env.sort_by = BYTES;
		} else if (!strcmp(arg, "time")) {
			env.sort_by = TIME;
		} else {
			warning("Invalid sort method: %s\n", arg);
			argp_usage(state);
		}
		break;
	case 'r':
		errno = 0;
		env.output_rows = strtol(arg, NULL, 10);
		if (errno || env.output_rows <= 0) {
			warning("Invalid rows: %s\n", arg);
			argp_usage(state);
		}
		if (env.output_rows > OUTPUT_ROWS_LIMIT)
			env.output_rows = OUTPUT_ROWS_LIMIT;
		break;
	case 'v':
		env.verbose = true;
		break;
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case ARGP_KEY_ARG:
		errno = 0;
		if (pos_args == 0) {
			env.interval = strtol(arg, NULL, 10);
			if (errno || env.interval <= 0) {
				warning("Invalid interval\n");
				argp_usage(state);
			}
		} else if (pos_args == 1) {
			env.count = strtol(arg, NULL, 10);
			if (errno || env.count <= 0) {
				warning("Invalid count\n");
				argp_usage(state);
			}
		} else {
			warning("Unrecognized positional argument: %s\n", arg);
			argp_usage(state);
		}
		pos_args++;
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}
	return 0;
}