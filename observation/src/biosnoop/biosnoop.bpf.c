// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "biosnoop.h"
#include "core_fixes.bpf.h"

#define MAX_ENTRIES	10240

const volatile bool filter_memcg = false;
const volatile bool target_queued = false;
const volatile bool filter_dev = false;
const volatile __u32 target_dev = 0;

extern __u32 LINUX_KERNEL_VERSION __kconfig;

struct {
	__uint(type, BPF_MAP_TYPE_CGROUP_ARRAY);
	__uint(max_entries, 1);
	__uint(key_size, sizeof(u32));
	__uint(value_size, sizeof(u32));
} cgroup_map SEC(".maps");

struct piddata {
	char comm[TASK_COMM_LEN];
	u32 pid;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct request *);
	__type(value, struct piddata);
} infobyreq SEC(".maps");

struct stage {
	u64 insert;
	u64 issue;
	__u32 dev;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct request *);
	__type(value, struct stage);
} start SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__type(key, u32);
	__type(value, u32);
} events SEC(".maps");

static __always_inline int trace_pid(struct request *rq)
{
	u64 id = bpf_get_current_pid_tgid();
	struct piddata piddata = {};

	piddata.pid = id >> 32;
	bpf_get_current_comm(&piddata.comm, sizeof(piddata.comm));
	bpf_map_update_elem(&infobyreq, &rq, &piddata, BPF_ANY);
	return 0;
}

SEC("fentry/blk_account_io_start")
int BPF_PROG(blk_account_io_start, struct request *rq)
{
	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	return trace_pid(rq);
}

SEC("kprobe/blk_account_io_start")
int BPF_KPROBE(kprobe_blk_account_io_start, struct request *rq)
{
	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	return trace_pid(rq);
}

SEC("kprobe/blk_account_io_merge_bio")
int BPF_KPROBE(blk_account_io_merge_bio, struct request *rq)
{
	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	return trace_pid(rq);
}

static __always_inline int trace_rq_start(struct request *rq, bool insert)
{
	struct stage *stagep, stage = {};
	u64 ts = bpf_ktime_get_ns();

	stagep = bpf_map_lookup_elem(&start, &rq);
	if (!stagep) {
		struct gendisk *disk = get_disk(rq);

		stage.dev = disk ? MKDEV(BPF_CORE_READ(disk, major),
					 BPF_CORE_READ(disk, first_minor)) : 0;
		if (filter_dev && target_dev != stage.dev)
			return 0;

		stagep = &stage;
	}

	if (insert)
		stagep->insert = ts;
	else
		stagep->issue = ts;

	if (stagep == &stage)
		bpf_map_update_elem(&start, &rq, stagep, BPF_ANY);
	return 0;
}

SEC("tp_btf/block_rq_insert")
int BPF_PROG(block_rq_insert)
{
	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	if (LINUX_KERNEL_VERSION >= KERNEL_VERSION(5, 11, 0))
		return trace_rq_start((void *)ctx[0], true);
	else
		return trace_rq_start((void *)ctx[1], true);
}

SEC("raw_tp/block_rq_insert")
int BPF_PROG(block_rq_insert_raw)
{
	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	if (LINUX_KERNEL_VERSION >= KERNEL_VERSION(5, 11, 0))
		return trace_rq_start((void *)ctx[0], true);
	else
		return trace_rq_start((void *)ctx[1], true);
}

