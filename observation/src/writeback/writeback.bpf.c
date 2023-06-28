// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
// Author: Copyright @ 2023 - Jackie Liu
//
// Based on bpftrace/writeback.bt - Brendan Gregg
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "compat.bpf.h"
#include "maps.bpf.h"
#include "writeback.h"

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, 1024);
	__type(key, dev_t);
	__type(value, __u64);
} birth SEC(".maps");

SEC("tracepoint/writeback/writeback_start")
int tracepoint_writeback_start(struct trace_event_raw_writeback_work_class *ctx)
{
	dev_t sb_dev = ctx->sb_dev;
	__u64 start = bpf_ktime_get_ns();

	bpf_map_update_elem(&birth, &sb_dev, &start, BPF_ANY);
	return 0;
}

