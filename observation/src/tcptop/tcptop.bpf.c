// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_endian.h>
#include "tcptop.h"
#include "maps.bpf.h"

/* Taken from kernel include/linux/socket.h */
#define AF_INET         2       /* Internet IP protocol */
#define AF_INET6        10      /* IP version 6 */

const volatile bool filter_cg = false;
const volatile pid_t target_pid = -1;
const volatile int target_family = -1;

struct {
        __uint(type, BPF_MAP_TYPE_CGROUP_ARRAY);
        __uint(max_entries, 1);
        __type(key, u32);
        __type(value, u32);
} cgroup_map SEC(".maps");

struct {
        __uint(type, BPF_MAP_TYPE_HASH);
        __uint(max_entries, 10240);
        __type(key, struct ip_key_t);
        __type(value, struct traffic_t);
} ip_map SEC(".maps");

