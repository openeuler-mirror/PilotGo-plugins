// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)
#include "vmlinux.h"
#include <bpf/bpf_core_read.h>
#include <bpf/bpf_tracing.h>
#include <bpf/bpf_helpers.h>
#include "ttysnoop.h"
#include "compat.bpf.h"
#include "core_fixes.bpf.h"

#define WRITE		1
#define ITER_UBUF	6

const volatile int user_data_count = 16;
const volatile int pts_inode = -1;

static int
do_tty_write(void *ctx, const struct file *file, const char *buf, size_t count)
{
	if (BPF_CORE_READ(file, f_inode, i_ino) != pts_inode)
		return 0;

	if (count < 0)
		return 0;

	for (int i = 0; i < user_data_count && count; i++) {
		struct event *event = reserve_buf(sizeof(*event));

		if (!event)
			break;

		 /**
		  * bpf_probe_read_user() can only use a fixed size, so truncate
		  * to count in user space
		  */
		if (bpf_probe_read_user(&event->buf, BUFSIZE, (void *)buf)) {
			discard_buf(event);
			break;
		}

		if (count > BUFSIZE) {
			event->buf[BUFSIZE] = 0;
			event->count = BUFSIZE;
		} else {
			event->count = count;
			event->buf[count] = 0;
		}

		submit_buf(ctx, event, sizeof(*event));

		if (count < BUFSIZE)
			break;

		count -= BUFSIZE;
		buf += BUFSIZE;
	}

	return 0;
}

