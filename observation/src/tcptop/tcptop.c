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

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
        switch (key) {
        case 'p':
                env.target_pid = argp_parse_pid(key, arg, state);
                break;
        case 'c':
                env.cgroup_path = arg;
                env.cgroup_filtering = true;
                break;
        case 'C':
                env.clear_screen = false;
                break;
        case 'S':
                env.no_summary = true;
                break;
        case '4':
                env.ipv4_only = true;
                break;
        case '6':
                env.ipv6_only = true;
                break;
        case 's':
                if (!strcmp(arg, "all")) {
                        env.sort_by = ALL;
                } else if (!strcmp(arg, "sent")) {
                        env.sort_by = SENT;
                } else if (!strcmp(arg, "received")) {
                        env.sort_by = RECEIVED;
                } else {
                        warning("Invalid sort method: %s\n", arg);
                        argp_usage(state);
                }
                break;
        case 'r':
                env.output_rows = MIN(argp_parse_long(key, arg, state), OUTPUT_ROWS_LIMIT);
                break;
        case 'v':
                env.verbose = true;
                break;
        case 'h':
                argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
                break;
        case ARGP_KEY_END:
                if (env.ipv4_only && env.ipv6_only) {
                        warning("Only one --ipvX option should be used\n");
                        argp_usage(state);
                }
                break;
        case ARGP_KEY_ARG:
                if (state->arg_num == 0) {
                        env.interval = argp_parse_long(key, arg, state);
                } else if (state->arg_num == 1) {
                        env.count = argp_parse_long(key, arg, state);
                } else {
                        warning("Unrecognized positional argument: %s\n", arg);
                        argp_usage(state);
                }
                break;
        default:
                return ARGP_ERR_UNKNOWN;
        }
        return 0;
}

static int libbpf_print_fn(enum libbpf_print_level level, const char *format,
                           va_list args)
{
        if (level == LIBBPF_DEBUG && !env.verbose)
                return 0;
        return vfprintf(stderr, format, args);
}

static void sig_handler(int sig)
{
        exiting = 1;
}

static int sort_column(const void *obj1, const void *obj2)
{
        struct info_t *i1 = (struct info_t *)obj1;
        struct info_t *i2 = (struct info_t *)obj2;

        if (i1->key.family != i2->key.family) {
                /*
                 * i1 - i2 because we want to sort by increasing order (first
                 * AF_INET then AF_INET6).
                 */
                return i1->key.family - i2->key.family;
        }

        if (env.sort_by == SENT)
                return i2->value.sent - i1->value.sent;
        else if (env.sort_by == RECEIVED)
                return i2->value.received - i1->value.received;
        else
                return (i2->value.sent + i2->value.received) - (i1->value.sent + i1->value.received);
}

