// SPDX-License-Identifier: GPL-2.0
#include "commons.h"
#include "btf_helpers.h"
#include "trace_helpers.h"
#include "tcplife.h"
#include "tcplife.skel.h"
#include "compat.h"

#include <arpa/inet.h>

static volatile bool exiting = false;

static struct env {
        pid_t   target_pid;
        short   target_family;
        __u16   target_sports[MAX_PORTS];
        bool    filter_sport;
        __u16   target_dports[MAX_PORTS];
        bool    filter_dport;
        int     column_width;
        bool    emit_timestamp;
        bool    verbose;
} env = {
        .column_width = 15,
};

const char *argp_program_version = "tcplife 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"Trace the lifespan of TCP sessions and summarize.\n"
"\n"
"USAGE: tcplife [-h] [-p PID] [-4] [-6] [-L] [-R] [-T] [-w]\n"
"\n"
"EXAMPLES:\n"
"    tcplife -p 1215             # only trace PID 1215\n"
"    tcplife -p 1215 -4          # trace IPv4 only\n";

static const struct argp_option opts[] = {
        { "pid", 'p', "PID", 0, "Process ID to trace" },
        { "ipv4", '4', NULL, 0, "Trace IPv4 only" },
        { "ipv6", '6', NULL, 0, "Trace IPv6 only" },
        { "wide", 'w', NULL, 0, "Wide column output (fits IPv6 addesses)" },
        { "time", 'T', NULL, 0, "Include timestamp on output" },
        { "localport", 'L', "LOCALPORT", 0, "Comma-separated list of local ports to trace." },
        { "remoteport", 'R', "REMOTEPORT", 0, "Comma-separated list of remote ports to trace." },
        { "verbose", 'v', NULL, 0, "Verbose debug output" },
        { NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
        {}
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
        switch (key) {
        case 'p':
                env.target_pid = argp_parse_pid(key, arg, state);
                break;
        case '4':
                env.target_family = AF_INET;
                break;
        case '6':
                env.target_family = AF_INET6;
                break;
        case 'w':
                env.column_width = 26;
                break;
        case 'L':
        case 'R':
        {
                char *port = strtok(arg, ",");

                for (int i = 0; i < MAX_PORTS && port; i++) {
                        if (key == 'L')
                                env.target_sports[i] = safe_strtol(port, 1, 65535, state);
                        else
                                env.target_dports[i] = safe_strtol(port, 1, 63355, state);
                        port = strtok(NULL, ",");
                }
                break;
        }
        case 'T':
                env.emit_timestamp = true;
                break;
        case 'v':
                env.verbose = true;
                break;
        case 'h':
                argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
                break;
        case ARGP_KEY_END:
                if (env.target_sports[0] != 0)
                        env.filter_sport = true;
                if (env.target_dports[0] != 0)
                        env.filter_dport = true;
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

static int print_events(struct bpf_buffer *buf)
{
        int err;

        err = bpf_buffer__open(buf, handle_event, handle_lost_events, NULL);
        if (err) {
                warning("Failed to open ring/perf buffer\n");
                return err;
        }

        if (env.emit_timestamp)
                printf("%-8s ", "TIME(s)");
        printf("%-7s %-16s %-*s %-5s %-*s %-5s %-6s %-6s %-s\n",
               "PID", "COMM", env.column_width, "LADDR", "LPORT",
               env.column_width, "RADDR", "RPORT", "TX_KB", "RX_KB", "MS");

        while (!exiting) {
                err = bpf_buffer__poll(buf, POLL_TIMEOUT_MS);
                if (err < 0 && err != -EINTR) {
                        warning("Error polling ring/perf buffer: %s\n", strerror(-err));
                        break;
                }
                /* reset err to return 0 if exiting */
                err = 0;
        }

        return err;
}

