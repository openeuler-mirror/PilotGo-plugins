// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include "tcptop.h"
#include "tcptop.skel.h"
#include "trace_helpers.h"

#include <arpa/inet.h>
#include <sys/param.h>

#define OUTPUT_ROWS_LIMIT       10240
#define IPV4                    0
#define PORT_LENGTH             5

enum SORT {
        ALL,
        SENT,
        RECEIVED,
};

static volatile sig_atomic_t exiting;

static struct env {
        pid_t target_pid;
        char *cgroup_path;
        bool cgroup_filtering;
        bool clear_screen;
        bool no_summary;
        bool ipv4_only;
        bool ipv6_only;
        int output_rows;
        int sort_by;
        int interval;
        int count;
        bool verbose;
} env = {
        .target_pid = -1,
        .clear_screen = true,
        .output_rows = 20,
        .interval = 1,
        .count = 99999999,
};

const char *argp_program_version = "tcptop 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace sending and received operation over IP.\n"
"\n"
"USAGE: tcptop [-h] [-p PID] [interval] [count]\n"
"\n"
"EXAMPLES:\n"
"    tcptop            # TCP top, refresh every 1s\n"
"    tcptop -p 1216    # only trace PID 1216\n"
"    tcptop -c path    # only trace the given cgroup path\n"
"    tcptop 5 10       # 5s summaries, 10 times\n";

static const struct argp_option opts[] = {
        { "pid", 'p', "PID", 0, "Process ID to trace" },
        { "cgroup", 'c', "/sys/fs/cgroup/unified", 0, "Trace process in cgroup path" },
        { "ipv4", '4', NULL, 0, "Trace IPv4 family only" },
        { "ipv6", '6', NULL, 0, "Trace IPv6 family only" },
        { "nosummary", 'S', NULL, 0, "Skip system summary line" },
        { "noclear", 'C', NULL, 0, "Don't clear the screen" },
        { "sort", 's', "SORT", 0, "Sort columns, default all [all, sent, received]" },
        { "rows", 'r', "ROW", 0, "Maximum rows to print, default 20" },
        { "verbose", 'v', NULL, 0, "Verbose debug output" },
        { NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
        {}
};

struct info_t {
        struct ip_key_t key;
        struct traffic_t value;
};

