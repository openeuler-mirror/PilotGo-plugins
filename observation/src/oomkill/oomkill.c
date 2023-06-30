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