// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "fsslower.h"
#include "bits.bpf.h"
#include "maps.bpf.h"

#define MAX_ENTRIES	8192

const volatile pid_t target_pid = 0;
const volatile __u64 min_lat_ns = 0;

struct data {
	__u64 ts;
	loff_t start;
	loff_t end;
	struct file *fp;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, __u32);
	__type(value, struct data);
} starts SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__uint(key_size, sizeof(u32));
	__uint(value_size, sizeof(u32));
} events SEC(".maps");

static int probe_entry(struct file *fp, loff_t start, loff_t end)
{
	__u64 pid_tgid = bpf_get_current_pid_tgid();
	__u32 pid = pid_tgid >> 32;
	__u32 tid = pid_tgid;
	struct data data;

	if (!fp)
		return 0;

	if (target_pid && target_pid != pid)
		return 0;

	data.ts = bpf_ktime_get_ns();
	data.start = start;
	data.end = end;
	data.fp = fp;
	bpf_map_update_elem(&starts, &tid, &data, BPF_ANY);
	return 0;
}
