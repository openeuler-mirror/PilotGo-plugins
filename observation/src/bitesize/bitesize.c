// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "bitesize.h"
#include "bitesize.skel.h"
#include "trace_helpers.h"

struct argument {
	char *disk;
	char *comm;
	int comm_len;
	time_t interval;
	bool timestamp;
	int times;
};

static volatile bool verbose = false;
static volatile sig_atomic_t exiting;

const char *argp_program_version = "bitesize 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Summarize block device I/O size as a histogram.\n"
"\n"
"USAGE: bitesize [--help] [-T] [-c COMM] [-d DISK] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    bitesize              # summarize block I/O latency as a histogram\n"
"    bitesize 1 10         # print 1 second summaries, 10 times\n"
"    bitesize -T 1         # 1s summaries with timestamps\n"
"    bitesize -c fio       # trace fio only\n";

static const struct argp_option opts[] = {
	{ "timestamp", 'T', NULL, 0, "Include timestamp on output" },
	{ "comm", 'c', "COMM", 0, "Trace this comm only" },
	{ "disk", 'd', "DISK", 0, "Trace this disk only" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	static int pos_args;
	struct argument *argument = state->input;
	size_t len = TASK_COMM_LEN;

	switch (key) {
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		verbose = true;
		break;
	case 'c':
		argument->comm = arg;
		argument->comm_len = min(strlen(arg) + 1, len);
		break;
	case 'd':
		argument->disk = arg;
		if (strlen(arg) + 1 > DISK_NAME_LEN) {
			warning("Invalid disk name: too long\n");
			argp_usage(state);
		}
		break;
	case 'T':
		argument->timestamp = true;
		break;
	case ARGP_KEY_ARG:
		errno = 0;
		if (pos_args == 0) {
			argument->interval = strtol(arg, NULL, 10);
			if (errno || argument->interval <= 0) {
				warning("Invalid interval\n");
				argp_usage(state);
			}
		} else if (pos_args == 1) {
			argument->times = strtol(arg, NULL, 10);
			if (errno || argument->times <= 0) {
				warning("Invalid times\n");
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

static int libbpf_print_fn(enum libbpf_print_level level, const char *format,
			   va_list args)
{
	if (level == LIBBPF_DEBUG && !verbose)
		return 0;
	return vfprintf(stderr, format, args);
}

static void sig_handler(int sig)
{
	exiting = 1;
}

static int print_log2_hists(int fd)
{
	struct hist_key lookup_key = {}, next_key;
	struct hist hist;
	int err;

	while (!bpf_map_get_next_key(fd, &lookup_key, &next_key)) {
		err = bpf_map_lookup_elem(fd, &next_key, &hist);
		if (err < 0) {
			warning("Failed to lookup hist: %d\n", err);
			return -1;
		}
		printf("\nProcess Name = %s\n", next_key.comm);
		print_log2_hist(hist.slots, MAX_SLOTS, "Kbytes");
		lookup_key = next_key;
	}

	memset(lookup_key.comm, '?', sizeof(lookup_key.comm));
	while (!bpf_map_get_next_key(fd, &lookup_key, &next_key)) {
		err = bpf_map_delete_elem(fd, &next_key);
		if (err < 0) {
			warning("Failed to cleanup hist : %d\n", err);
			return -1;
		}
		lookup_key = next_key;
	}

	return 0;
}