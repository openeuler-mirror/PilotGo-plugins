// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_helpers.h>
#include <bpf/bpf_tracing.h>
#include "capable.h"

#define MAX_ENTRIES	10240

extern int LINUX_KERNEL_VERSION	__kconfig;

struct myinfo myinfo = {};

const volatile enum uniqueness unique_type = UNQ_OFF;
const volatile bool kernel_stack = false;
const volatile bool user_stack = false;
const volatile bool filter_cg = false;
const volatile pid_t target_pid = -1;

struct args_t {
	int cap;
	int cap_out;
};

struct unique_key {
	int cap;
	u32 tgid;
	u64 cgroupid;
};

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, u64);
	__type(value, struct args_t);
} start SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_CGROUP_ARRAY);
	__uint(max_entries, 1);
	__type(key, u32);
	__type(value, u32);
} cgroup_map SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__type(key, __u32);
	__type(value, __u32);
} events SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_STACK_TRACE);
	__type(key, u32);
} stackmap SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct key_t);
	__type(value, struct cap_event);
} info SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, MAX_ENTRIES);
	__type(key, struct unique_key);
	__type(value, u64);
} seen SEC(".maps");

SEC("kprobe/cap_capable")
int BPF_KPROBE(kprobe__cap_capable_entry, const struct cred *cred,
	       struct user_namespace *target_ns, int cap, int cap_out)
{
	__u64 pid_tgid;
	struct bpf_pidns_info nsdata;

	if (filter_cg && !bpf_current_task_under_cgroup(&cgroup_map, 0))
		return 0;


	if (bpf_get_ns_current_pid_tgid(myinfo.dev, myinfo.ino, &nsdata,
					sizeof(struct bpf_pidns_info)))
		return 0;

	pid_tgid = (__u64)nsdata.tgid << 32 | nsdata.pid;
	if (myinfo.pid_tgid == pid_tgid)
		return 0;

	if (target_pid != -1 && target_pid != nsdata.tgid)
		return 0;

	struct args_t args = {};
	args.cap = cap;
	args.cap_out = cap_out;

	bpf_map_update_elem(&start, &pid_tgid, &args, BPF_ANY);

	return 0;
}
