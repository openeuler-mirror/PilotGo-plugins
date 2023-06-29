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

#define OPT_PERF_MAX_STACK_DEPTH 1 /* --perf-max-stack-depth */
#define OPT_STACK_STORAGE_SIZE 2   /* --stack-storage-size */
#define OPT_STATE 3                /* --state */

static const struct argp_option opts[] = {
    {"pid", 'p', "PID", 0, "Trace this PID only"},
    {"tid", 't', "TID", 0, "Trace this TID only"},
    {"user-threads-only", 'u', NULL, 0,
     "User threads only (no kernel threads)"},
    {"kernel-threads-only", 'k', NULL, 0,
     "Kernel threads only (no user threads)"},
    {"perf-max-stack-depth", OPT_PERF_MAX_STACK_DEPTH,
     "PERF-MAX-STACK-DEPTH", 0, "the limit for both kernel and user stack traces (default 127)"},
    {"stack-storage-size", OPT_STACK_STORAGE_SIZE, "STACK-STORAGE-SIZE", 0,
     "the number of unique stack traces that can be stored and displayed (default 1024)"},
    {"min-block-time", 'm', "MIN-BLOCK-TIME", 0,
     "the amount of time in microseconds over which we store traces (default 1)"},
    {"max-block-time", 'M', "MAX-BLOCK-TIME", 0,
     "the amount of time in microseconds under which we store traces (default U64_MAX)"},
    {"state", OPT_STATE, "STATE", 0,
     "filter on this thread state bitmask (eg, 2 == TASK_UNINTERRUPTIBLE) see include/linux/sched.h"},
    {"verbose", 'v', NULL, 0, "Verbose debug output"},
    {NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help"},
    {},
};