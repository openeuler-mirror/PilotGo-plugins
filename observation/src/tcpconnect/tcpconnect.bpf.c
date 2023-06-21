// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>

#include "tcpconnect.h"
#include "compat.bpf.h"
#include "maps.bpf.h"

const volatile int filter_ports[MAX_PORTS] = {};
const volatile int filter_ports_len = 0;
const volatile uid_t filter_uid = -1;
const volatile pid_t filter_pid = 0;
const volatile bool do_count = false;
const volatile bool source_port = false;

/* Define here, because there are conflicts with include files */
#define AF_INET         2
#define AF_INET6        10

struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, MAX_ENTRIES);
        __type(key, u32);
        __type(value, struct sock *);
} sockets SEC(".maps");

struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, MAX_ENTRIES);
        __type(key, struct ipv4_flow_key);
        __type(value, u64);
} ipv4_count SEC(".maps");

struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, MAX_ENTRIES);
        __type(key, struct ipv6_flow_key);
        __type(value, u64);
} ipv6_count SEC(".maps");

static __always_inline bool filter_port(__u16 port)
{
        if (filter_ports_len == 0)
                return false;

        for (int i = 0; i < filter_ports_len && i < MAX_PORTS; i++) {
                if (port == filter_ports[i])
                        return false;
        }

        return true;
}

static __always_inline void
count_v4(struct sock *sk, __u16 sport, __u16 dport)
{
        struct ipv4_flow_key key = {};
        static __u64 zero;
        __u64 *val;

        BPF_CORE_READ_INTO(&key.saddr, sk, __sk_common.skc_rcv_saddr);
        BPF_CORE_READ_INTO(&key.daddr, sk, __sk_common.skc_daddr);
        key.sport = sport;
        key.dport = dport;
        val = bpf_map_lookup_or_try_init(&ipv4_count, &key, &zero);
        if (!val)
                return;
        __atomic_add_fetch(val, 1, __ATOMIC_RELAXED);
}

static __always_inline void
count_v6(struct sock *sk, __u16 sport, __u16 dport)
{
        struct ipv6_flow_key key = {};
        static const __u64 zero;
        __u64 *val;

        BPF_CORE_READ_INTO(&key.saddr, sk,
                           __sk_common.skc_v6_rcv_saddr.in6_u.u6_addr32);
        BPF_CORE_READ_INTO(&key.daddr, sk,
                           __sk_common.skc_v6_daddr.in6_u.u6_addr32);
        key.sport = sport;
        key.dport = dport;

        val = bpf_map_lookup_or_try_init(&ipv6_count, &key, &zero);
        if (!val)
                return;
        __atomic_add_fetch(val, 1, __ATOMIC_RELAXED);
}

static __always_inline void
trace_v4(void *ctx, pid_t pid, struct sock *sk, __u16 sport, __u16 dport)
{
        struct event *event;

        event = reserve_buf(sizeof(*event));
        if (!event)
                return;

        event->af = AF_INET;
        event->pid = pid;
        event->uid = bpf_get_current_uid_gid();
        BPF_CORE_READ_INTO(&event->saddr_v4, sk, __sk_common.skc_rcv_saddr);
        BPF_CORE_READ_INTO(&event->daddr_v4, sk, __sk_common.skc_daddr);
        event->sport = sport;
        event->dport = dport;
        bpf_get_current_comm(&event->task, sizeof(event->task));

        submit_buf(ctx, event, sizeof(*event));
}

static __always_inline void
trace_v6(void *ctx, pid_t pid, struct sock *sk, __u16 sport, __u16 dport)
{
        struct event *event;

        event = reserve_buf(sizeof(*event));
        if (!event)
                return;

        event->af = AF_INET6;
        event->pid = pid;
        event->uid = bpf_get_current_uid_gid();
        BPF_CORE_READ_INTO(&event->saddr_v6, sk,
                           __sk_common.skc_v6_rcv_saddr.in6_u.u6_addr32);
        BPF_CORE_READ_INTO(&event->daddr_v6, sk,
                           __sk_common.skc_v6_daddr.in6_u.u6_addr32);
        event->sport = sport;
        event->dport = dport;
        bpf_get_current_comm(&event->task, sizeof(event->task));

        submit_buf(ctx, event, sizeof(*event));
}

