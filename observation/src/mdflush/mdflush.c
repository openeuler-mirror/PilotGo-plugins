// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "mdflush.h"
#include "mdflush.skel.h"
#include "trace_helpers.h"

static volatile sig_atomic_t exiting;
static volatile bool verbose = false;

const char *argp_program_version = "mdflush 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
    "Trace md flush events.\n"
    "\n"
    "USAGE: mdflush\n";

static const struct argp_option opts[] = {
    {"verbose", 'v', NULL, 0, "Verbose debug output"},
    {NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help"},
    {}};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
    switch (key)
    {
    case 'v':
        verbose = true;
        break;
    case 'h':
        argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
        break;
    default:
        return ARGP_ERR_UNKNOWN;
    }
    return 0;
}

static void sig_handler(int sig)
{
    exiting = 1;
}

static int libbpf_print_fn(enum libbpf_print_level level, const char *format,
                           va_list args)
{
    if (level == LIBBPF_DEBUG && !verbose)
        return 0;
    return vfprintf(stderr, format, args);
}

static void handle_event(void *ctx, int cpu, void *data, __u32 data_sz)
{
    struct event *e = data;
    char ts[32];

    strftime_now(ts, sizeof(ts), "%H%M:%S");
    printf("%-8s %-7d %-16s %-s\n", ts, e->pid, e->comm, e->disk);
}

static void handle_lost_events(void *ctx, int cpu, __u64 lost_cnt)
{
    warning("Lost %llu events on CPU #%d\n", lost_cnt, cpu);
}
