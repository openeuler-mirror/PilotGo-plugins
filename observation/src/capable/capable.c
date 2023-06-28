// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "capable.h"
#include "capable.skel.h"
#include "trace_helpers.h"
#include <sys/stat.h>
#include <sys/syscall.h>

struct argument {
	char *cgroupspath;
	bool cg;
	bool extra_fields;
	bool user_stack;
	bool kernel_stack;
	bool unique;
	char *unique_type;
	int stack_storage_size;
	int perf_max_stack_depth;
	pid_t pid;
};

struct context {
	struct argument *argument;
	int ifd;
	int sfd;
};

const char *cap_name[] = {
	[0] = "CAP_CHOWN",
	[1] = "CAP_DAC_OVERRIDE",
	[2] = "CAP_DAC_READ_SEARCH",
	[3] = "CAP_FOWNER",
	[4] = "CAP_FSETID",
	[5] = "CAP_KILL",
	[6] = "CAP_SETGID",
	[7] = "CAP_SETUID",
	[8] = "CAP_SETPCAP",
	[9] = "CAP_LINUX_IMMUTABLE",
	[10] = "CAP_NET_BIND_SERVICE",
	[11] = "CAP_NET_BROADCAST",
	[12] = "CAP_NET_ADMIN",
	[13] = "CAP_NET_RAW",
	[14] = "CAP_IPC_LOCK",
	[15] = "CAP_IPC_OWNER",
	[16] = "CAP_SYS_MODULE",
	[17] = "CAP_SYS_RAWIO",
	[18] = "CAP_SYS_CHROOT",
	[19] = "CAP_SYS_PTRACE",
	[20] = "CAP_SYS_PACCT",
	[21] = "CAP_SYS_ADMIN",
	[22] = "CAP_SYS_BOOT",
	[23] = "CAP_SYS_NICE",
	[24] = "CAP_SYS_RESOURCE",
	[25] = "CAP_SYS_TIME",
	[26] = "CAP_SYS_TTY_CONFIG",
	[27] = "CAP_MKNOD",
	[28] = "CAP_LEASE",
	[29] = "CAP_AUDIT_WRITE",
	[30] = "CAP_AUDIT_CONTROL",
	[31] = "CAP_SETFCAP",
	[32] = "CAP_MAC_OVERRIDE",
	[33] = "CAP_MAC_ADMIN",
	[34] = "CAP_SYSLOG",
	[35] = "CAP_WAKE_ALARM",
	[36] = "CAP_BLOCK_SUSPEND",
	[37] = "CAP_AUDIT_READ",
	[38] = "CAP_PERFMON",
	[39] = "CAP_BPF",
	[40] = "CAP_CHECKPOINT_RESTORE"
};

static volatile sig_atomic_t exiting;
static volatile bool verbose = false;
struct syms_cache *syms_cache = NULL;
struct ksyms *ksyms = NULL;
int ifd, sfd;

const char *argp_program_version = "capable 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace security capability checks (cap_capable()).\n"
"\n"
"USAGE: capable [--help] [-p PID | -c CG | -K | -U | -x] [-u TYPE]\n"
"[--perf-max-stack-depth] [--stack-storage-size]\n"
"\n"
"EXAMPLES:\n"
"    capable                  # Trace capability checks\n"
"    capable -p 185           # Trace this PID only\n"
"    capable -c CG            # Trace process under cgroupsPath CG\n"
"    capable -K               # Add kernel stacks to trace\n"
"    capable -x               # Extra fields: show TID and INSETID columns\n"
"    capable -U               # Add user-space stacks to trace\n"
"    capable -u TYPE          # Print unique output for TYPE=[pid | cgroup] (default:off)\n";

#define OPT_PERF_MAX_STACK_DEPTH	1	/* for --perf-max-stack-depth */
#define OPT_STACK_STORAGE_SIZE		2	/* for --stack-storage-size */

static const struct argp_option opts[] = {
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ "pid", 'p', "PID", 0, "Trace this PID only" },
	{ "cgroup", 'c', "/sys/fs/cgroup/unifed", 0, "Trace process in cgroup path" },
	{ "kernel-stack", 'K', NULL, 0, "output kernel stack trace" },
	{ "user-stack", 'U', NULL, 0, "output user stack trace" },
	{ "extra-fields", 'x', NULL, 0, "extra fields: show TID and INSETID columns" },
	{ "unique", 'u', "off", 0, "Print unique output for <pid> or <cgroup> (default: off)" },
	{ "perf-max-stack-depth", OPT_PERF_MAX_STACK_DEPTH,
	  "PERF-MAX-STACK-DEPTH", 0, "the limit for both kernel and user stack traces (default: 127)" },
	{ "stack-storage-size", OPT_STACK_STORAGE_SIZE, "STACK-STORAGE-SIZE", 0,
	  "the number of unique stack traces that can be stored and displayed (default: 1024)" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

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
