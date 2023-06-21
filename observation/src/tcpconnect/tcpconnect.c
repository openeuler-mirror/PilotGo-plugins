// SPDX-License-Identifier: GPL-2.0
#include "commons.h"
#include "tcpconnect.h"
#include "tcpconnect.skel.h"
#include "btf_helpers.h"
#include "trace_helpers.h"
#include "compat.h"
#include "map_helpers.h"
#include <arpa/inet.h>

static struct timespec start_time;
static volatile bool exiting = false;

const char *argp_program_version = "tcpconnect 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
"\ntcpconnect: Count/Trace active tcp connections\n"
"\n"
"EXAMPLES:\n"
"    tcpconnect             # trace all TCP connect()s\n"
"    tcpconnect -t          # include timestamps\n"
"    tcpconnect -p 181      # only trace PID 181\n"
"    tcpconnect -P 80       # only trace port 80\n"
"    tcpconnect -P 80,81    # only trace port 80 and 81\n"
"    tcpconnect -U          # include UID\n"
"    tcpconnect -u 1000     # only trace UID 1000\n"
"    tcpconnect -c          # count connects per src, dest, port\n"
"    tcpconnect --C mappath # only trace cgroups in the map\n"
"    tcpconnect --M mappath # only trace mount namespaces in the map\n"
;

static const struct argp_option opts[] = {
        { "verbose", 'v', NULL, 0, "Verbose debug output" },
        { "timestamp", 't', NULL, 0, "Include timestamp on output" },
        { "count", 'c', NULL, 0, "Count connects per src ip and dst ip/port" },
        { "print-uid", 'U', NULL, 0, "Include UID on output" },
        { "pid", 'p', "PID", 0, "Process PID to trace" },
        { "uid", 'u', "UID", 0, "Process UID to trace" },
        { "source-port", 's', NULL, 0, "Consider source port when counting" },
        { "port", 'P', "PORTS", 0,
          "Comma-separated list of destination ports to trace" },
        { "cgroupmap", 'C', "PATH", 0, "Trace cgroups in this map" },
        { "mntnsmap", 'M', "PATH", 0, "Trace mount namespaces in this map" },
        { NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help" },
        {}
};

static struct env {
        bool verbose;
        bool count;
        bool print_timestamp;
        bool print_uid;
        pid_t pid;
        uid_t uid;
        int nports;
        int ports[MAX_PORTS];
        bool source_port;
} env = {
        .uid = (uid_t)-1
};

static error_t parse_arg(int key, char *arg, struct argp_state *state)
{
        switch (key) {
        case 'h':
                argp_state_help(state, stderr, ARGP_HELP_STD_HELP);
                break;
        case 'v':
                env.verbose = true;
                break;
        case 'c':
                env.count = true;
                break;
        case 's':
                env.source_port = true;
                break;
        case 't':
                env.print_timestamp = true;
                break;
        case 'U':
                env.print_uid = true;
                break;
        case 'p':
                env.pid = argp_parse_pid(key, arg, state);
                break;
        case 'u':
                env.uid = safe_strtoul(arg, 0, (uid_t)-2, state);
                break;
        case 'P':
        {
                char *port = strtok(arg, ",");

                for (int i = 0; port; i++) {
                        env.ports[i] = safe_strtol(port, 1, 65535, state);
                        env.nports++;

                        port = strtok(NULL, ",");
                }
                break;
        }
        case 'C':
                warning("Not implemented: --cgroupmap");
                break;
        case 'M':
                warning("Not implemented: --mntnsmap");
                break;
        default:
                return ARGP_ERR_UNKNOWN;
        }

        return 0;
}

