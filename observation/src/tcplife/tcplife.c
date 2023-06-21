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

