// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package str

import "sort"

type Set map[string]struct{}

func MakeSet(strings ...string) Set {
	if len(strings) == 0 {
		return nil
	}

	set := Set{}
	for _, str := range strings {
		set[str] = struct{}{}
	}
	return set
}

func (set Set) Add(s string) {
	set[s] = struct{}{}
}

func (set Set) Del(s string) {
	delete(set, s)
}

func (set Set) Count() int {
	return len(set)
}

func (set Set) Has(s string) (exists bool) {
	if set != nil {
		_, exists = set[s]
	}
	return exists
}

// Equals compares this StringSet with another StringSet.
func (set Set) Equals(anotherSet Set) bool {
	if set.Count() != anotherSet.Count() {
		return false
	}

	for k := range set {
		if !anotherSet.Has(k) {
			return false
		}
	}

	return true
}

// ToSlice returns the items in the set as a sorted slice.
func (set Set) ToSlice() []string {
	keys := make([]string, 0, len(set))
	for key := range set {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	return keys
}
