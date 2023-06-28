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