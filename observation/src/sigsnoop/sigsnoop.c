// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "sigsnoop.h"
#include "sigsnoop.skel.h"
#include "compat.h"

#include <libgen.h>

static volatile sig_atomic_t exiting;

static pid_t target_pid = 0;
static int target_signal = 0;
static bool failed_only = false;
static bool kill_only = false;
static bool signal_name = false;
static bool verbose = false;

const char *argp_program_version = "sigsnoop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace standard and real-time signals.\n"
"\n"
"USAGE: sigsnoop [-h] [-x] [-k] [-n] [-p PID] [-s SIGNAL]\n"
"\n"
"EXAMPLES:\n"
"    sigsnoop             # trace signals system-wide\n"
"    sigsnoop -k          # trace signals issued by kill syscall only\n"
"    sigsnoop -x          # trace failed signals only\n"
"    sigsnoop -p 1216     # only trace PID 1216\n"
"    sigsnoop -s 9        # only trace signal 9\n";

static const struct argp_option opts[] = {
	{ "failed", 'x', NULL, 0, "Trace failed signals only." },
	{ "kill", 'k', NULL, 0, "Trace signals issued by kill syscall only." },
	{ "pid", 'p', "PID", 0, "Process ID to trace" },
	{ "signal", 's', "SIGNAL", 0, "Signal to trace." },
	{ "name", 'n', NULL, 0, "Output signal name instead of signal number." },
	{ "verbose", 'v', NULL, 0, "verbose debug output" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	switch (key) {
	case 'p':
		target_pid = argp_parse_pid(key, arg, state);
		break;
	case 's':
		target_signal = argp_parse_long(key, arg, state);
		break;
	case 'n':
		signal_name = true;
		break;
	case 'x':
		failed_only = true;
		break;
	case 'k':
		kill_only = true;
		break;
	case 'v':
		verbose = true;
		break;
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
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

static void alias_parse(char *prog)
{
	char *name = basename(prog);

	if (!strcmp(name, "killsnoop"))
		kill_only = true;
}

int main(int argc, char *argv[])
{
	static const struct argp argp = {
		.options = opts,
		.parser = parse_arg,
		.doc = argp_program_doc,
	};
	struct bpf_buffer *buf = NULL;
	struct sigsnoop_bpf *obj;
	int err;

	alias_parse(argv[0]);
	err = argp_parse(&argp, argc, argv, 0, NULL, NULL);
	if (err)
		return err;
        
	if (!bpf_is_root())
		return 1;

	libbpf_set_print(libbpf_print_fn);	
        
        obj = sigsnoop_bpf__open();
	if (!obj) {
		warning("Failed to open BPF object\n");
		return 1;
	}	

        obj->rodata->filtered_pid = target_pid;
	obj->rodata->target_signal = target_signal;
	obj->rodata->failed_only = failed_only;

	if (kill_only) {
		bpf_program__set_autoload(obj->progs.sig_trace, false);
	} else {
		bpf_program__set_autoload(obj->progs.kill_entry, false);
		bpf_program__set_autoload(obj->progs.kill_exit, false);
		bpf_program__set_autoload(obj->progs.tkill_entry, false);
		bpf_program__set_autoload(obj->progs.tkill_exit, false);
		bpf_program__set_autoload(obj->progs.tgkill_entry, false);
		bpf_program__set_autoload(obj->progs.tgkill_exit, false);
	}

        buf = bpf_buffer__new(obj->maps.events, obj->maps.heap);
	if (!buf) {
		err = -errno;
		warning("Failed to create ring/perf buffer: %d\n", err);
		goto cleanup;
	}

	err = sigsnoop_bpf__load(obj);
	if (err) {
		warning("Failed to load BPF object: %d\n", err);
		goto cleanup;
	}

	err = sigsnoop_bpf__attach(obj);
	if (err) {
		warning("Failed to attach BPF programs: %d\n", err);
		goto cleanup;
	}

	err = bpf_buffer__open(buf, handle_event, handle_lost_events, NULL);
	if (err) {
		warning("Failed to open ring/perf buffer: %d\n", err);
		goto cleanup;
	}

	return err != 0;
}

