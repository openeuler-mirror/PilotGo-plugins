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

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	struct argument *argument = state->input;

	switch (key) {
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		verbose = true;
		break;
	case 'p':
		argument->pid = argp_parse_pid(key, arg, state);
		break;
	case 'c':
		argument->cgroupspath = arg;
		argument->cg = true;
		break;
	case 'U':
		argument->user_stack = true;
		break;
	case 'K':
		argument->kernel_stack = true;
		break;
	case 'x':
		argument->extra_fields = true;
		break;
	case 'u':
		argument->unique_type = arg;
		argument->unique = true;
		break;
	case OPT_PERF_MAX_STACK_DEPTH:
		errno = 0;
		argument->perf_max_stack_depth = strtol(arg, NULL, 10);
		if (errno || argument->perf_max_stack_depth == 0) {
			warning("Invalid perf max stack depth: %s\n", arg);
			argp_usage(state);
		}
		break;
	case OPT_STACK_STORAGE_SIZE:
		errno = 0;
		argument->stack_storage_size = strtol(arg, NULL, 10);
		if (errno || argument->stack_storage_size <= 0) {
			warning("Invalid stack storage size: %s\n", arg);
			argp_usage(state);
		}
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

static void print_map(struct ksyms *ksyms, struct syms_cache *syms_cache,
		      void *ctx)
{
	struct key_t lookup_key = {}, next_key;
	const struct ksym *ksym;
	const struct syms *syms;
	const struct sym *sym;
	int err;
	unsigned long *ip;
	struct cap_event val;
	int ifd = ((struct context *)ctx)->ifd;
	int sfd = ((struct context *)ctx)->sfd;
	struct argument *argument = ((struct context *)ctx)->argument;

	ip = calloc(argument->perf_max_stack_depth, sizeof(*ip));
	if (!ip) {
		warning("Failed to alloc ip\n");
		return;
	}

	while (!bpf_map_get_next_key(ifd, &lookup_key, &next_key)) {
		err = bpf_map_lookup_elem(ifd, &next_key, &val);
		if (err < 0) {
			warning("Failed to lookup info: %d\n", err);
			goto cleanup;
		}
		lookup_key = next_key;

		if (argument->kernel_stack) {
			if (bpf_map_lookup_elem(sfd, &next_key.kernel_stack_id, ip))
				warning("    [Missed Kernel Stack]\n");
			for (int i = 0; i < argument->perf_max_stack_depth && ip[i]; i++) {
				ksym = ksyms__map_addr(ksyms, ip[i]);
				printf("    %s\n", ksym ? ksym->name : "Unknown");
			}
		}

		if (argument->user_stack) {
			if (next_key.user_stack_id == -1)
				goto skip_ustack;

			if (bpf_map_lookup_elem(sfd, &next_key.user_stack_id, ip)) {
				warning("    [Missed User Stack]\n");
				continue;
			}

			syms = syms_cache__get_syms(syms_cache, next_key.tgid);
			if (!syms) {
				warning("Failed to get syms\n");
				goto skip_ustack;
			}

			for (int i = 0; i < argument->perf_max_stack_depth && ip[i]; i++) {
				sym = syms__map_addr(syms, ip[i]);
				printf("    %s\n", sym ? sym->name : "Unknown");
			}
		}

skip_ustack:
		printf("    %-16s %s (%d)\n", "-", val.task, next_key.pid);
	}

cleanup:
	free(ip);
}

static void handle_event(void *ctx, int cpu, void *data, __u32 data_sz)
{
	const struct cap_event *e = data;
	char ts[32];

	strftime_now(ts, sizeof(ts), "%H:%M:%S");

	char *verdict = "deny";
	if (!e->ret)
		verdict = "allow";

	if (((struct context *)ctx)->argument->extra_fields)
		printf("%-8s %-5d %-7d %-7d %-16s %-7d %-20s %-7d %-7s %-7d\n",
		       ts, e->uid, e->pid, e->tgid, e->task, e->cap, cap_name[e->cap],
		       e->audit, verdict, e->insetid);
	else
		printf("%-8s %-5d %-7d %-16s %-7d %-20s %-7d %-7s\n", ts,
		       e->uid, e->pid, e->task, e->cap, cap_name[e->cap],
		       e->audit, verdict);

	print_map(ksyms, syms_cache, ctx);
}

static void handle_lost_events(void *ctx, int cpu, __u64 lost_cnt)
{
	warning("Lost %llu events on CPU #%d!\n", lost_cnt, cpu);
}