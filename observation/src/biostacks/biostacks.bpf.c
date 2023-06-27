// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "biostacks.h"
#include "maps.bpf.h"
#include "core_fixes.bpf.h"
#include "bits.bpf.h"

#define MAX_ENTRIES	10240

const volatile bool target_ms = false;
const volatile bool filter_dev = false;
const volatile __u32 target_dev = 0;

static __always_inline int trace_start(void *ctx, struct request *rq, bool merge_bio)
{
	struct internal_rqinfo *i_rqinfop = NULL, i_rqinfo = {};
	struct gendisk *disk = get_disk(rq);
	u32 dev;

	dev = disk ? MKDEV(BPF_CORE_READ(disk, major),
			   BPF_CORE_READ(disk, first_minor)) : 0;
	if (filter_dev && dev != target_dev)
		return 0;

	if (merge_bio)
		i_rqinfop = bpf_map_lookup_elem(&rqinfos, &rq);
	if (!i_rqinfop)
		i_rqinfop = &i_rqinfo;

	i_rqinfop->start_ts = bpf_ktime_get_ns();
	i_rqinfop->rqinfo.pid = bpf_get_current_pid_tgid();
	i_rqinfop->rqinfo.kern_stack_size =
		bpf_get_stack(ctx, i_rqinfop->rqinfo.kern_stack,
			      sizeof(i_rqinfop->rqinfo.kern_stack), 0);
	bpf_get_current_comm(&i_rqinfop->rqinfo.comm,
			     sizeof(i_rqinfop->rqinfo.comm));
	i_rqinfop->rqinfo.dev = dev;

	if (i_rqinfop == &i_rqinfo)
		bpf_map_update_elem(&rqinfos, &rq, i_rqinfop, BPF_ANY);

	return 0;
}