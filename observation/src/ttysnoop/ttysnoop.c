// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "ttysnoop.h"
#include "ttysnoop.skel.h"
#include "compat.h"
#include <sys/stat.h>
#include <bpf/btf.h>
#include <sys/utsname.h>
#include "btf_helpers.h"

static volatile sig_atomic_t exiting;

static struct env {
	bool verbose;
	bool clear_screen;
	int count;
	int pts_inode;
	bool record;
	char *record_filename;
} env = {
	.clear_screen = true,
	.pts_inode = -1,
	.count = 16,
};

const char *argp_program_version = "ttysnoop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Watch live output from a tty or pts device.\n"
"\n"
"USAGE:   ttysnoop [-Ch] {PTS | /dev/ttydev}  # try -h for help\n"
"\n"
"Example:\n"
"    ttysnoop /dev/pts/2          # snoop output from /dev/pts/2\n"
"    ttysnoop 2                   # snoop output from /dev/pts/2 (shortcut)\n"
"    ttysnoop /dev/console        # snoop output from the system console\n"
"    ttysnoop /dev/tty0           # snoop output from /dev/tty0\n"
"    ttysnoop /dev/pts/2 -c 2     # snoop output from /dev/pts/2 with 2 checks\n"
"                                   for 256 bytes of data in buffer\n"
"                                   (potentially retrieving 512 bytes)\n";

static const struct argp_option opts[] = {
	{ "verbose", 'v', NULL, 0, "Verbose debug output" },
	{ "noclear", 'C', NULL, 0, "Don't clear the screen" },
	{ "datacount", 'c', "COUNT", 0, "Number of times we check for 'data-size' data (default 16)" },
	{ "record", 'r', "RECORD", 0, "Record tty history" },
	{ NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
	{}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
	switch (key) {
	case 'h':
		argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
		break;
	case 'v':
		env.verbose = true;
		break;
	case 'C':
		env.clear_screen = false;
		break;
	case 'c':
		env.count = argp_parse_long(key, arg, state);
		break;
	case 'r':
		env.record = true;
		env.record_filename = arg;
		break;
	case ARGP_KEY_ARG:
		if (state->arg_num != 0) {
			warning("Unrecognized positional arguments: %s\n", arg);
			argp_usage(state);
		}

		char path[4096] = {};
		struct stat st;

		if (arg[0] != '/') {
			strcpy(path, "/dev/pts/");
			strcat(path, arg);
		} else {
			strcpy(path, arg);
		}

		if (stat(path, &st)) {
			warning("Failed to stat console file: %s\n", arg);
			argp_usage(state);
		}
		env.pts_inode = st.st_ino;
		break;
	case ARGP_KEY_END:
		if (env.pts_inode == -1)
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
	if (level == LIBBPF_DEBUG && !env.verbose)
		return 0;
	return vfprintf(stderr, format, args);
}

static bool tty_write_is_newly(void)
{
	const struct btf_type *type;
	__s32 id;
	struct btf *btf;

	btf = btf__load_vmlinux_btf();
	if (!btf) {
		warning("No BTF, cannot determine type info: %s", strerror(errno));
		goto failed;
	}

	id = btf__find_by_name_kind(btf, "tty_write", BTF_KIND_FUNC);
	if (id <= 0) {
		warning("Can't find function tty_write in BTF: %s\n",
			strerror(-id));
		goto failed;
	}

	type = btf__type_by_id(btf, id);
	if (!type || BTF_INFO_KIND(type->info) != BTF_KIND_FUNC)
		goto failed;

	type = btf__type_by_id(btf, type->type);
	if (!type || BTF_INFO_KIND(type->info) != BTF_KIND_FUNC_PROTO)
		goto failed;

	btf__free(btf);
	/* the newly tty_write has 2 params, old have 4 params */
	if (btf_vlen(type) != 2)
		return false;
	return true;

failed:
	return fallback_to_compare_kernel_version();
}

int main(int argc, char *argv[])
{
	LIBBPF_OPTS(bpf_object_open_opts, open_opts);
	const struct argp argp = {
		.options = opts,
		.parser = parse_arg,
		.doc = argp_program_doc,
	};
	struct ttysnoop_bpf *obj;
	int err;

	err = argp_parse(&argp, argc, argv, 0, NULL, NULL);
	if (err)
		return err;
        
	if (!bpf_is_root())
		return 1;

	err = ensure_core_btf(&open_opts);
	if (err) {
		warning("Failed to fetch necessary BTF for CO-RE: %s\n", strerror(-err));
		return 1;
	}

	new_tty_write = tty_write_is_newly();
	libbpf_set_print(libbpf_print_fn);
	
	obj = ttysnoop_bpf__open_opts(&open_opts);
	if (!obj) {
		warning("Failed to open BPF object\n");
		return 1;
	}

        obj->rodata->user_data_count = env.count;
	obj->rodata->pts_inode = env.pts_inode;

return err != 0;
}
