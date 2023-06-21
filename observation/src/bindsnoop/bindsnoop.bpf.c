// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_endian.h>
#include "bindsnoop.h"

#define MAX_ENTRIES	10240
#define MAX_PORTS	1024

const volatile bool filter_memcg = false;
const volatile pid_t target_pid = 0;
const volatile bool ignore_errors = true;
const volatile bool filter_by_port = false;

static int probe_entry(struct pt_regs *ctx, struct socket *socket)
{
	__u64 pid_tgid = bpf_get_current_pid_tgid();
	pid_t tgid = pid_tgid >> 32;
	pid_t pid = (pid_t)pid_tgid;

	if (target_pid && target_pid != tgid)
		return 0;

	bpf_map_update_elem(&sockets, &pid, &socket, BPF_ANY);
	return 0;
}