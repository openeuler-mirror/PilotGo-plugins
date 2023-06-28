// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>

#include "maps.bpf.h"
#include "core_fixes.bpf.h"
#include "memleak.h"

const volatile size_t min_size = 0;
const volatile size_t max_size = -1;
const volatile size_t page_size = 4096;
const volatile __u64 sample_rate = 1;
const volatile bool trace_all = false;
const volatile __u64 stack_flags = 0;
const volatile bool wa_missing_free = false;

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 10240);
    __type(key, pid_t);
    __type(value, u64);
} sizes SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, ALLOCS_MAX_ENTRIES);
    __type(key, u64);
    __type(value, struct alloc_info);
} allocs SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, COMBINED_ALLOCS_MAX_ENTRIES);
    __type(key, u64);
    __type(value, union combined_alloc_info);
} combined_allocs SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_HASH);
    __uint(max_entries, 10240);
    __type(key, u64);
    __type(value, u64);
} memptrs SEC(".maps");

struct
{
    __uint(type, BPF_MAP_TYPE_STACK_TRACE);
    __type(key, u32);
} stack_traces SEC(".maps");

static union combined_alloc_info initial_cinfo;

static __always_inline void update_statistics_add(u64 stack_id, u64 sz)
{
    union combined_alloc_info *existing_cinfo;

    existing_cinfo = bpf_map_lookup_or_try_init(&combined_allocs, &stack_id,
                                                &initial_cinfo);
    if (!existing_cinfo)
        return;

    const union combined_alloc_info incremental_cinfo = {
        .total_size = sz,
        .number_of_allocs = 1};

    __sync_fetch_and_add(&existing_cinfo->bits, incremental_cinfo.bits);
}

static __always_inline void update_statistics_del(u64 stack_id, u64 sz)
{
    union combined_alloc_info *existing_cinfo;

    existing_cinfo = bpf_map_lookup_elem(&combined_allocs, &stack_id);
    if (!existing_cinfo)
    {
        bpf_printk("Failed to lookup combined allocs\n");
        return;
    }

    const union combined_alloc_info decremental_cinfo = {
        .total_size = sz,
        .number_of_allocs = 1,
    };

    __sync_fetch_and_add(&existing_cinfo->bits, -decremental_cinfo.bits);
}

static __always_inline int gen_alloc_enter(size_t size)
{
    if (size < min_size || size > max_size)
        return 0;

    if (sample_rate > 1)
    {
        if (bpf_ktime_get_ns() % sample_rate != 0)
            return 0;
    }

    const pid_t pid = bpf_get_current_pid_tgid() >> 32;
    bpf_map_update_elem(&sizes, &pid, &size, BPF_ANY);

    if (trace_all)
        bpf_printk("alloc entered, size = %lu\n", size);

    return 0;
}