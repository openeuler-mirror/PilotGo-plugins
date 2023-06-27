// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "commons.h"
#include <numa.h>
#include "numasched.skel.h"
#include "numasched.h"
#include "trace_helpers.h"

static volatile sig_atomic_t exiting;

static struct env
{
    bool verbose;
    bool timestamp;
    char *comm;
    pid_t pid;
    pid_t tid;
} env = {
    .pid = INVALID_PID,
    .tid = INVALID_PID,
};

const char *argp_program_version = "numasched 0.1";
const char *argp_program_bug_address = "Jackie Liu <liuyun01@kylinos.cn>";
const char argp_program_doc[] =
    "Trace task NUMA switch\n"
    "\n"
    "USAGE: numasched [-p PID] [-t TID] [-c COMM]\n"
    "\n"
    "EXAMPLES:\n"
    "    ./numasched             # Trace all numa node switch\n"
    "    ./numasched -p 123      # Trace pid 123 only\n"
    "    ./numasched -t 1234     # Trace thread id 1234 only\n"
    "    ./numasched -c comm     # Trace this comm only\n"
    "    ./numasched -T          # Include timestamp\n"
    "    ./numasched -v          # Verbose debug output\n";

static const struct argp_option opts[] = {
    {"timestamp", 'T', NULL, 0, "Include timestamp"},
    {"verbose", 'v', NULL, 0, "Verbose debug output"},
    {"pid", 'p', "PID", 0, "Trace this PID only"},
    {"tid", 't', "TID", 0, "Trace this TID only"},
    {"comm", 'c', "COMM", 0, "Trace this comm only"},
    {NULL, 'h', NULL, OPTION_HIDDEN, "Show the full help"},
    {},
};