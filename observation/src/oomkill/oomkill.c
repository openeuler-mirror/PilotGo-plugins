// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "oomkill.h"
#include "oomkill.skel.h"
#include "compat.h"
#include "btf_helpers.h"
#include "trace_helpers.h"

static volatile sig_atomic_t exiting;
static bool verbose = false;

const char *argp_program_version = "oomkill 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
    "Trace OOM kills.\n"
    "\n"
    "USAGE: oomkill [-h]\n"
    "\n"
    "EXAMPLES:\n"
    "    oomkill               # trace OOM kills\n";

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

static int handle_event(void *ctx, void *data, size_t len)
{
    FILE *f;
    char buf[256];
    int n = 0;
    char ts[32];
    struct data_t *e = data;

    f = fopen("/proc/loadavg", "r");
    if (f)
    {
        memset(buf, 0, sizeof(buf));
        n = fread(buf, 1, sizeof(buf), f);
        fclose(f);
    }
    strftime_now(ts, sizeof(ts), "%H:%M:%S");

    printf("%s Trigger by PID %d (\"%s\"), OOM kill of PID %d (\"%s\"), %lld pages",
           ts, e->fpid, e->fcomm, e->tpid, e->tcomm, e->pages);
    if (n)
        printf(", loadavg: %s", buf);
    else
        printf("\n");

    return 0;
}
static void handle_lost_events(void *ctx, int cpu, __u64 lost_cnt)
{
    warning("Lost %llu events on CPU #%d!\n", lost_cnt, cpu);
}

static void sig_handler(int sig)
{
    exiting = 1;
}