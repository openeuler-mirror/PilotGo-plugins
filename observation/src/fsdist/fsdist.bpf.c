// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "bits.bpf.h"
#include "fsdist.h"

#define MAX_ENTRIES	10240

const volatile pid_t target_pid = 0;
const volatile bool in_ms = false;

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, __u32);
	__type(value, __u64);
} starts SEC(".maps");

struct hist hists[F_MAX_OP] = {};

static int probe_entry()
{
	__u64 pid_tgid = bpf_get_current_pid_tgid();
	__u32 pid = pid_tgid >> 32;
	__u32 tid = (__u32)pid_tgid;
	__u64 ts;

	if (target_pid && target_pid != pid)
		return 0;

	ts = bpf_ktime_get_ns();
	bpf_map_update_elem(&starts, &tid, &ts, BPF_ANY);
	return 0;
}
