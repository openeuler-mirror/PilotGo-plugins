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

static __always_inline int gen_alloc_exit2(void *ctx, u64 address)
{
    const pid_t pid = bpf_get_current_pid_tgid() >> 32;
    struct alloc_info info = {};

    const u64 *size = bpf_map_lookup_and_delete_elem(&sizes, &pid);
    if (!size)
        return 0;

    info.size = *size;

    if (address != 0)
    {
        info.timestamp_ns = bpf_ktime_get_ns();
        info.stack_id = bpf_get_stackid(ctx, &stack_traces, stack_flags);
        bpf_map_update_elem(&allocs, &address, &info, BPF_ANY);
        update_statistics_add(info.stack_id, info.size);
    }

    if (trace_all)
        bpf_printk("alloc exited, size = %lu, result = %lx\n",
                   info.size, address);

    return 0;
}

static __always_inline int gen_alloc_exit(struct pt_regs *ctx)
{
    return gen_alloc_exit2(ctx, PT_REGS_RC(ctx));
}

static __always_inline int gen_free_enter(const void *address)
{
    const u64 addr = (u64)address;

    const struct alloc_info *info = bpf_map_lookup_and_delete_elem(&allocs, &addr);
    if (!info)
        return 0;

    update_statistics_del(info->stack_id, info->size);

    if (trace_all)
        bpf_printk("Free entered, address = %lx, size = %lu\n",
                   address, info->size);

    return 0;
}

SEC("uprobe")
int BPF_KPROBE(malloc_enter, size_t size)
{
    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(malloc_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(free_enter, void *address)
{
    return gen_free_enter(address);
}

SEC("uprobe")
int BPF_KPROBE(calloc_enter, size_t nmemb, size_t size)
{
    return gen_alloc_enter(nmemb * size);
}

SEC("uretprobe")
int BPF_KRETPROBE(calloc_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(realloc_enter, void *ptr, size_t size)
{
    gen_free_enter(ptr);

    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(realloc_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(mmap_enter, void *address, size_t size)
{
    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(mmap_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(munmap_enter, void *address)
{
    return gen_free_enter(address);
}

SEC("uprobe")
int BPF_KPROBE(posix_memalign_enter, void **memptr, size_t alignment, size_t size)
{
    const __u64 memptr64 = (__u64)(size_t)memptr;
    const __u64 pid = bpf_get_current_pid_tgid() >> 32;

    bpf_map_update_elem(&memptrs, &pid, &memptr64, BPF_ANY);

    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(posix_memalign_exit)
{
    const __u64 pid = bpf_get_current_pid_tgid() >> 32;
    __u64 *memptr64;

    memptr64 = bpf_map_lookup_and_delete_elem(&memptrs, &pid);
    if (!memptr64)
        return 0;

    return gen_alloc_exit2(ctx, *memptr64);
}

SEC("uprobe")
int BPF_KPROBE(aligned_alloc_enter, size_t alignment, size_t size)
{
    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(aligned_alloc_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(valloc_enter, size_t size)
{
    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(valloc_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(memalign_enter, size_t alignment, size_t size)
{
    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(memalign_exit)
{
    return gen_alloc_exit(ctx);
}

SEC("uprobe")
int BPF_KPROBE(pvalloc_enter, size_t size)
{
    return gen_alloc_enter(size);
}

SEC("uretprobe")
int BPF_KRETPROBE(pvalloc_exit)
{
    return gen_alloc_exit(ctx);
}