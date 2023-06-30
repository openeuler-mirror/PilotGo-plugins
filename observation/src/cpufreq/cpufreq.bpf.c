// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include "cpufreq.h"
#include "maps.bpf.h"

__u32 freqs_mhz[MAX_CPU_NR] = {};
static struct hist zero;
struct hist syswide = {};
bool filter_memcg = false;

struct {
	__uint(type, BPF_MAP_TYPE_CGROUP_ARRAY);
	__type(key, u32);
	__type(value, u32);
	__uint(max_entries, 1);
} cgroup_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct hkey);
	__type(value, struct hist);
} hists SEC(".maps");

static __always_inline int probe_cpu_frequency(unsigned int state, unsigned int cpu_id)
{
	if (filter_memcg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;

	if (cpu_id >= MAX_CPU_NR)
		return 0;

	freqs_mhz[cpu_id] = state / 1000;
	return 0;
}