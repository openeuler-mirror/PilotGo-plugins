// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "biosnoop.h"
#include "core_fixes.bpf.h"

const volatile bool filter_memcg = false;
const volatile bool target_queued = false;
const volatile bool filter_dev = false;
const volatile __u32 target_dev = 0;

extern __u32 LINUX_KERNEL_VERSION __kconfig;

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

static __always_inline int trace_pid(struct request *rq)
{
	u64 id = bpf_get_current_pid_tgid();
	struct piddata piddata = {};

	piddata.pid = id >> 32;
	bpf_get_current_comm(&piddata.comm, sizeof(piddata.comm));
	bpf_map_update_elem(&infobyreq, &rq, &piddata, BPF_ANY);
	return 0;
}