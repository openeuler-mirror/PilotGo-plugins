// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>

#include "biolatency.h"
#include "bits.bpf.h"
#include "core_fixes.bpf.h"
#include "maps.bpf.h"

#define MAX_ENTRIES	10240

extern __u32 LINUX_KERNEL_VERSION __kconfig;

const volatile bool filter_memcg = false;
const volatile bool target_per_disk = false;
const volatile bool target_per_flag = false;
const volatile bool target_queued = false;
const volatile bool target_ms = false;
const volatile bool filter_dev = false;
const volatile __u32 target_dev = 0;

static int __always_inline trace_rq_start(struct request *rq, int issue)
{
	u64 ts;

	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	if (issue && target_queued && BPF_CORE_READ(rq, q, elevator))
		return 0;

	ts = bpf_ktime_get_ns();

	if (filter_dev) {
		struct gendisk *disk = get_disk(rq);
		u32 dev;

		dev = disk ? MKDEV(BPF_CORE_READ(disk, major),
				BPF_CORE_READ(disk, first_minor)) : 0;
		if (target_dev != dev)
			return 0;
	}

	bpf_map_update_elem(&start, &rq, &ts, BPF_ANY);

	return 0;
}