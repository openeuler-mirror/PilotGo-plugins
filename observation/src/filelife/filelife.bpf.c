// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "filelife.h"
#include "core_fixes.bpf.h"

/* linux: include/linux/fs.h */
#define FMODE_CREATED	0x100000

const volatile pid_t target_tgid = 0;

struct {
	__uint(type, BPF_MAP_TYPE_HASH);
	__uint(max_entries, 8192);
	__type(key, struct dentry *);
	__type(value, u64);
} start SEC(".maps");

struct {
	__uint(type, BPF_MAP_TYPE_PERF_EVENT_ARRAY);
	__uint(key_size, sizeof(u32));
	__uint(value_size, sizeof(u32));
} events SEC(".maps");

static __always_inline int probe_create(struct dentry *dentry)
{
	u64 id = bpf_get_current_pid_tgid();
	u32 tgid = id >> 32;
	u64 ts;

	if (target_tgid && target_tgid != tgid)
		return 0;

	ts = bpf_ktime_get_ns();
	bpf_map_update_elem(&start, &dentry, &ts, BPF_ANY);

	return 0;
}

/*
 * In different kernel versions, function vfs_create() has two declarations,
 * and their parameter lists are as follows:
 *
 * int vfs_create(struct inode *dir, struct dentry *dentry, umode_t mode,
 *		  bool want_excl);
 * int vfs_create(struct user_namespace *mnt_userns, struct inode *dir,
 *		  struct dentry *dentry, umode_t mode, bool want_excl);
 */
SEC("kprobe/vfs_create")
int BPF_KPROBE(vfs_create, void *arg0, void *arg1, void *arg2)
{
	if (renamedata_has_old_mnt_userns_field())
		return probe_create(arg2);
	else
		return probe_create(arg1);
}
