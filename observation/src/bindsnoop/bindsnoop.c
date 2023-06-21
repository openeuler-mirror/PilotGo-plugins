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
const char argp_program_doc[] =
"Trace bind syscalls.\n"
"\n"
"USAGE: bindsnoop [-h] [-t] [-x] [-p PID] [-P ports] [-c CG]\n"
"\n"
"EXAMPLES:\n"
"    bindsnoop             # trace all bind syscall\n"
"    bindsnoop -t          # include timestamps\n"
"    bindsnoop -x          # include errors on output\n"
"    bindsnoop -p 1216     # only trace PID 1216\n"
"    bindsnoop -c CG       # Trace process under cgroupsPath CG\n"
"    bindsnoop -P 80,81    # only trace port 80 and 81\n"
"\n"
"Socket options are reported as:\n"
"  SOL_IP     IP_FREEBIND              F....\n"
"  SOL_IP     IP_TRANSPARENT           .T...\n"
"  SOL_IP     IP_BIND_ADDRESS_NO_PORT  ..N..\n"
"  SOL_SOCKET SO_REUSEADDR             ...R.\n"
"  SOL_SOCKET SO_REUSEPORT             ....r\n";

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

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	char *port;

	switch (key) {
	case 'p':
		env.target_pid = argp_parse_pid(key, arg, state);
		break;
	case 'c':
		env.cgroupspath = arg;
		env.cg = true;
		break;
	case 'P':
		if (!arg) {
			warning("No ports specified\n");
			argp_usage(state);
		}
		env.target_ports = strdup(arg);
		port = strtok(arg, ",");
		while (port) {
			int port_num = strtol(port, NULL, 10);
			if (errno || port_num <= 0 || port_num > 65536) {
				warning("Invalid ports: %s\n", arg);
				argp_usage(state);
			}
			port = strtok(NULL, ",");
		}
		break;
	case 'x':
		env.ignore_errors = false;
		break;
	case 't':
		env.emit_timestamp = true;
		break;
	case 'v':
		env.verbose = true;
		break;
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	default:
		return ARGP_ERR_UNKNOWN;
	}

	return 0;
}
