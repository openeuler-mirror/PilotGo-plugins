// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "memleak.h"
#include "memleak.skel.h"
#include "trace_helpers.h"

#ifdef USE_BLAZESYM
#include "blazesym.h"
#endif

#include <sys/eventfd.h>
#include <sys/wait.h>
#include <sys/param.h>

#define DEFAULT_MIN_AGE_NS 500

static struct env
{
    int interval;
    int nr_intervals;
    pid_t pid;
    bool pid_from_child;
    bool trace_all;
    bool show_allocs;
    bool combined_only;
    int64_t min_age_ns;
    uint64_t sample_rate;
    int top_stacks;
    size_t min_size;
    size_t max_size;
    char object[32];

    bool wa_missing_free;
    bool percpu;
    int perf_max_stack_depth;
    int stack_map_max_entries;
    long page_size;
    bool kernel_trace;
    bool verbose;
    char command[32];
} env = {
    .interval = 5,
    .nr_intervals = -1,
    .pid = -1,
    .min_age_ns = DEFAULT_MIN_AGE_NS,
    .sample_rate = 1,
    .top_stacks = 10,
    .max_size = -1,
    .perf_max_stack_depth = 127,
    .stack_map_max_entries = 10240,
    .page_size = -1,
    .kernel_trace = true,
};