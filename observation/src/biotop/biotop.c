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