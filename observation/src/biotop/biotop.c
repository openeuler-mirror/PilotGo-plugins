// SPDX-License-Identifier: (LGPL-2.1 OR BSD-2-Clause)

#include "commons.h"
#include "biotop.h"
#include "biotop.skel.h"
#include "trace_helpers.h"
#include "compat.h"

#define OUTPUT_ROWS_LIMIT	10240

enum SORT {
	ALL,
	IO,
	BYTES,
	TIME,
};

struct disk {
	int major;
	int minor;
	char name[256];
};

struct vector {
	size_t nr;
	size_t capacity;
	void **elems;
};

int grow_vector(struct vector *vector)
{
	if (vector->nr >= vector->capacity) {
		void **reallocated;

		if (!vector->capacity)
			vector->capacity = 1;
		else
			vector->capacity *= 2;

		reallocated = libbpf_reallocarray(vector->elems, vector->capacity, sizeof(*vector->elems));
		if (!reallocated)
			return -1;

		vector->elems = reallocated;
	}

	return 0;
}

void free_vector(struct vector vector)
{
	for (size_t i = 0; i < vector.nr; i++) {
		if (vector.elems[i] != NULL)
			free(vector.elems[i]);
	}

	free(vector.elems);
}

struct vector disks = {};

static volatile sig_atomic_t exiting;

static struct env {
	bool	clear_screen;
	int	output_rows;
	int	sort_by;
	int	interval;
	int	count;
	bool	verbose;
} env = {
	.clear_screen	= true,
	.output_rows	= 20,
	.sort_by	= ALL,
	.interval	= 1,
	.count		= 99999999,
	.verbose	= false,
};