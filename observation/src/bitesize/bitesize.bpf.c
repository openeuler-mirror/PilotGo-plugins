// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "bitesize.h"
#include "bits.bpf.h"
#include "maps.bpf.h"
#include "core_fixes.bpf.h"

const volatile char target_comm[TASK_COMM_LEN] = {};
const volatile bool filter_dev = false;
const volatile __u32 target_dev = 0;

extern __u32 LINUX_KERNEL_VERSION __kconfig;

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, 10240);
	__type(key, struct hist_key);
	__type(value, struct hist);
} hists SEC(".maps");

static struct hist zero;

static __always_inline bool comm_allowed(const char *comm)
{
	for (int i = 0; target_comm[i] != '\0' && i < TASK_COMM_LEN; i++) {
		if (comm[i] != target_comm[i])
			return false;
	}

	return true;
}
