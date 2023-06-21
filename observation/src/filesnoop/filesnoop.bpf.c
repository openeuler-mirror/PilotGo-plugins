// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
// Copyright @ 2023 - Kylin
// Author: Jackie Liu <liuyun01@kylinos.cn>

#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "filesnoop.h"
#include "compat.bpf.h"
#include "maps.bpf.h"

const volatile __u64 target_filename_sz = 0;
const volatile bool filter_filename = false;
const volatile int target_op = F_ALL;

#define MAX_ENTRIES	1024

char target_filename[FSFILENAME_MAX] = {};
