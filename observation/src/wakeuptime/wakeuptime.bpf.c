// SPDX-License-Identifier: GPL-3.0
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include "wakeuptime.h"
#include "maps.bpf.h"

#define PF_KTHREAD	0x00200000 /* kernel thread */

const volatile pid_t target_pid = 0;
const volatile __u64 max_block_ns = -1;
const volatile __u64 min_block_ns = 1;
const volatile bool user_threads_only = false;

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct key_t);
	__type(value, u64);
} counts SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, u32);
	__type(value, u64);
} start SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_STACK_TRACE);
	__uint(key_size, sizeof(u32));
} stackmap SEC(".maps");

int BPF_PROG(sched_switch_btf, bool preempt, struct task_struct *prev,
	     struct task_struct *next)
{
	return offcpu_sched_switch(prev);
}

SEC("raw_tp/sched_switch")
int BPF_PROG(sched_switch_raw, bool preempt, struct task_struct *prev,
	     struct task_struct *next)
{
	return offcpu_sched_switch(prev);
}

SEC("tp_btf/sched_wakeup")
int BPF_PROG(sched_wakeup_btf, struct task_struct *p)
{
	return wakeup(ctx, p);
}

SEC("raw_tp/sched_wakeup")
int BPF_PROG(sched_wakeup_raw, struct task_struct *p)
{
	return wakeup(ctx, p);
}

char LICENSE[] SEC("license") = "GPL";
