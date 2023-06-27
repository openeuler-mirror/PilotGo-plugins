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