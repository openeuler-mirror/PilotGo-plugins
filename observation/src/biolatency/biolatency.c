// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "biolatency.h"
#include "biolatency.skel.h"
#include "trace_helpers.h"
#include "blk_types.h"
#include <sys/resource.h>

static struct env {
	char	*disk;
	time_t	interval;
	int	times;
	bool	timestamp;
	bool	queued;
	bool	per_disk;
	bool	per_flag;
	bool	milliseconds;
	bool	verbose;
	char	*cgroupspath;
	bool	cg;
} env = {
	.interval = 99999999,
	.times = 99999999,
};

static volatile sig_atomic_t exiting;

const char *argp_program_version = "biolatency 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";