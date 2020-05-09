// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hashset

type HashSet struct {
	data map[string]struct{}
}

type HashUint32Set struct {
	data map[uint32]struct{}
}

func NewHashSet() HashSet {
	hs := HashSet{}
	hs.data = make(map[string]struct{})
	return hs
}

func NewHashSetPtr() *HashSet {
	hs := NewHashSet()
	return &hs
}

// clears the map without extra allocations
// faster since go 1.11
func (hs *HashSet) ClearFast() {
	for k := range hs.data {
		delete(hs.data, k)
	}
}

func NewHashUint32Set() HashUint32Set {
	hs := HashUint32Set{}
	hs.data = make(map[uint32]struct{})
	return hs
}

func NewHashUint32SetPtr() *HashUint32Set {
	hs := NewHashUint32Set()
	return &hs
}

// clears the map without extra allocations
// faster since go 1.11
func (hs *HashUint32Set) ClearFast() {
	for k := range hs.data {
		delete(hs.data, k)
	}
}
