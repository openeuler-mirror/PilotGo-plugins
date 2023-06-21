// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include "biotop.h"
#include "maps.bpf.h"
#include "core_fixes.bpf.h"

#define MAX_ENTRIES	10240

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct request *);
	__type(value, struct start_req_t);
} start SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct request *);
	__type(value, struct who_t);
} whobyreq SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct info_t);
	__type(value, struct val_t);
} counts SEC(".maps");

SEC("kprobe")
int BPF_KPROBE(blk_account_io_start, struct request *req)
{
	struct who_t who = {};

	/* cache PID and comm by request */
	bpf_get_current_comm(&who.name, sizeof(who.name));
	who.pid = bpf_get_current_pid_tgid() >> 32;
	bpf_map_update_elem(&whobyreq, &req, &who, BPF_ANY);

	return 0;
}

char LICENSE[] SEC("license") = "GPL";
