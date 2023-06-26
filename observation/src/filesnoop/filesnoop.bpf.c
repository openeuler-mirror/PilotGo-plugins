// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
// Copyright @ 2023 - Kylin
// Author: Jackie Liu <liuyun01@kylinos.cn>

#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "filesnoop.h"
#include "compat.bpf.h"
#include "maps.bpf.h"

const volatile __u64 target_filename_sz = 0;
const volatile bool filter_filename = false;
const volatile int target_op = F_ALL;

#define MAX_ENTRIES	1024

char target_filename[FSFILENAME_MAX] = {};

struct key_t {
	pid_t tid;
	int   fd;
};

struct fsfilename {
	char name[FSFILENAME_MAX];
};

struct print_value {
	struct key_t key;
	struct fsfilename *filename;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct key_t);
	__type(value, struct fsfilename);
} files SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, pid_t);
	__type(value, struct fsfilename);
} opens SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, pid_t);
	__type(value, struct print_value);
} prints SEC(".maps");

/* Filter filename */
static __always_inline bool filename_matched(const char *filename)
{
	if (!filter_filename)
		return true;

	for (int i = 0; i < target_filename_sz && i < FSFILENAME_MAX ; i++) {
		if (filename[i] != target_filename[i])
			return false;
	}

	return true;
}

/* Filter target operation */
static __always_inline bool is_target_operation(enum file_op op)
{
	switch (target_op) {
	case F_READ:
	case F_READV:
		return op == F_READ || op == F_READV;
	case F_WRITE:
	case F_WRITEV:
		return op == F_WRITE || op == F_WRITEV;
	case F_OPEN:
	case F_OPENAT:
	case F_OPENAT2:
		return op == F_OPEN || op == F_OPENAT || op == F_OPENAT2;
	case F_STATX:
	case F_FSTATFS:
	case F_NEWFSTAT:
		return op == F_STATX || op == F_FSTATFS || op == F_NEWFSTAT;
	case F_RENAMEAT:
	case F_RENAMEAT2:
		return op == F_RENAMEAT || op == F_RENAMEAT2;
	case F_UNLINKAT:
		return op == F_UNLINKAT;
	case F_CLOSE:
		return op == F_CLOSE;
	case F_UTIMENSAT:
		return op == F_UTIMENSAT;
	}

	return true;
}

static __always_inline int
handle_file_syscall_open_enter(struct trace_event_raw_sys_enter *ctx, enum file_op op)
{
	struct fsfilename filename = {};

	if (filter_filename && target_filename_sz == 0)
		return 0;

	pid_t tid = bpf_get_current_pid_tgid();

	if (op == F_OPENAT || op == F_OPENAT2)
		bpf_probe_read_user_str(&filename.name, FSFILENAME_MAX, (const char *)ctx->args[1]);
	else
		bpf_probe_read_user_str(&filename.name, FSFILENAME_MAX, (const char *)ctx->args[0]);

	/* If not match name, everything is over */
	if (!filename_matched(filename.name))
		return 0;

	bpf_map_update_elem(&opens, &tid, &filename, BPF_ANY);
	return 0;
}
