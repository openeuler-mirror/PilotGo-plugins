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
