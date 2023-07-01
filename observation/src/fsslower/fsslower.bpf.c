// SPDX-License-Identifier: GPL-2.0
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "fsslower.h"
#include "bits.bpf.h"
#include "maps.bpf.h"

#define MAX_ENTRIES	8192

const volatile pid_t target_pid = 0;
const volatile __u64 min_lat_ns = 0;

struct data {
	__u64 ts;
	loff_t start;
	loff_t end;
	struct file *fp;
};
