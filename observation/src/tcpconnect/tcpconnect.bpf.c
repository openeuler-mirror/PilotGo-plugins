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

