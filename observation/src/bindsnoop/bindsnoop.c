// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "bindsnoop.h"
#include "bindsnoop.skel.h"
#include "trace_helpers.h"
#include "btf_helpers.h"

#include <sys/socket.h>
#include <arpa/inet.h>
static struct env {
	char	*cgroupspath;
	bool	cg;
	bool	emit_timestamp;
	pid_t	target_pid;
	bool	ignore_errors;
	char	*target_ports;
	bool	verbose;
} env = {
	.ignore_errors = true,
};

static volatile sig_atomic_t exiting;

const char *argp_program_version = "bindsnoop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";

static const struct argp_option opts[] = {
	{ "timestamp", 't', NULL, 0, "Include timestamp on output" },
	{ "cgroup", 'c', "/sys/fs/cgroup/unified", 0, "Trace process in cgroup path" },
	{ "failed", 'x', NULL, 0, "Include errors on outputs" },
	{ "pid", 'p', "PID", 0, "Process ID to trace" },
	{ "ports", 'P', "PORTS", 0, "Comma-separated list of ports to trace" },
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};