// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_core_read.h>
#include "offcputime.h"
#include "core_fixes.bpf.h"

#define PF_KTHREAD 0x00200000 /* Kernel thread */
#define MAX_ENTRIES 10240

const volatile bool kernel_threads_only = false;
const volatile bool user_threads_only = false;
const volatile __u64 max_block_ns = -1;
const volatile __u64 min_block_ns = -1;
const volatile pid_t target_tgid = -1;
const volatile pid_t target_pid = -1;
const volatile long state = -1;

struct internal_key
{
    u64 start_ts;
    offcpu_key_t key;
};

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, u32);
    __type(value, struct internal_key);
    __uint(max_entries, MAX_ENTRIES);
} start SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_STACK_TRACE);
    __uint(key_size, sizeof(u32));
} stackmap SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __type(key, offcpu_key_t);
    __type(value, offcpu_val_t);
    __uint(max_entries, MAX_ENTRIES);
} info SEC(".maps");

static bool allow_record(struct task_struct *task)
{
    if (target_tgid != -1 && target_tgid != BPF_CORE_READ(task, tgid))
        return false;
    if (target_pid != -1 && target_pid != BPF_CORE_READ(task, pid))
        return false;
    if (user_threads_only && BPF_CORE_READ(task, flags) & PF_KTHREAD)
        return false;
    else if (kernel_threads_only && !(BPF_CORE_READ(task, flags) & PF_KTHREAD))
        return false;
    if (state != -1 && get_task_state(task) != state)
        return false;
    return true;
}