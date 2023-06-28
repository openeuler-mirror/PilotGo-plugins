// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "offcputime.h"
#include "offcputime.skel.h"
#include "trace_helpers.h"

static struct env
{
    pid_t pid;
    pid_t tid;
    bool user_threads_only;
    bool kernel_threads_only;
    int stack_storage_size;
    int perf_max_stack_depth;
    __u64 min_block_time;
    __u64 max_block_time;
    long state;
    int duration;
    bool verbose;
} env = {
    .pid = -1,
    .tid = -1,
    .stack_storage_size = 1024,
    .perf_max_stack_depth = 127,
    .min_block_time = 1,
    .max_block_time = -1,
    .state = -1,
    .duration = 99999999,
};

const char *argp_program_version = "offcputime 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
    "Summarize off-CPU time by stack trace.\n"
    "\n"
    "USAGE: offcputime [--help] [-p PID | -u | -k] [-m MIN-BLOCK-TIME] "
    "[-M MAX-BLOCK-TIME] [--state] [--perf-max-stack-depth] [--stack-storage-size] "
    "[duration]\n\n"
    "EXAMPLES:\n"
    "    offcputime             # trace off-CPU stack time until Ctrl-C\n"
    "    offcputime 5           # trace for 5 seconds only\n"
    "    offcputime -m 1000     # trace only events that last more than 1000 usec\n"
    "    offcputime -M 10000    # trace only events that last less than 10000 usec\n"
    "    offcputime -p 185      # only trace threads for PID 185\n"
    "    offcputime -t 188      # only trace thread 188\n"
    "    offcputime -u          # only trace user threads (no kernel)\n"
    "    offcputime -k          # only trace kernel threads (no user)\n";