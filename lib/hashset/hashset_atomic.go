// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hashset

import (
	"encoding/json"
	"io/ioutil"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/zerjioang/gotools/lib/hash/aeshash"
	"github.com/zerjioang/gotools/lib/logger"
	"github.com/zerjioang/gotools/util/str"
)

type HashSetAtomic struct {
	m  atomic.Value
	mu *sync.Mutex // used only by writers
}

func NewAtomicHashSet() HashSetAtomic {
	m := HashSetAtomic{}
	m.mu = new(sync.Mutex)
	return m
}

func NewAtomicHashSetPtr() HashSetAtomic {
	hs := NewAtomicHashSet()
	return hs
}

func (set *HashSetAtomic) Add(item string) {
	m1 := set.Read()         // load current value of the data structure
	m2 := NewHashUint32Set() // create a new value
	for k, v := range m1.data {
		m2.data[k] = v // copy all data from the current object to the new one
	}
	// hash item before adding it
	hashedKey := aeshash.Hash(item)
	set.mu.Lock()             // synchronize with other potential writers
	m2.data[hashedKey] = none // do the update that we need
	set.m.Store(m2)           // atomically replace the current object with the new one
	set.mu.Unlock()
	// At this point all new readers start working with the new version.
	// The old version will be garbage collected once the existing readers
	// (if any) are done with it.
}

// this call is considered as unsafe because write locks are removed from it
func (set *HashSetAtomic) UnsafeAddUint32(v uint32) {
	d := set.Read()
	d.data[v] = none // do the update that we need
	set.m.Store(d)   // atomically replace the current object with the new one
}

func (set *HashSetAtomic) Read() HashUint32Set {
	d := set.m.Load()
	var item HashUint32Set
	if d == nil {
		item = NewHashUint32Set()
		set.m.Store(item)
	} else {
		item = d.(HashUint32Set)
	}
	return item
}

func (set *HashSetAtomic) Size() int {
	return len(set.Read().data)
}

func (set *HashSetAtomic) Clear() {
	set.mu.Lock() // synchronize with other potential writers
	set.m.Store(NewHashUint32Set())
	set.mu.Unlock()
}

func (set *HashSetAtomic) ClearFast() {
	set.mu.Lock() // synchronize with other potential writers
	hs := set.Read()
	hs.ClearFast()
	set.mu.Unlock()
}

func (set *HashSetAtomic) ContainsString(item string) bool {
	source := set.m.Load().(HashUint32Set)
	// hash item
	hashedKey := aeshash.Hash(item)
	_, contains := source.data[hashedKey]
	return contains
}

func (set *HashSetAtomic) Contains(item uint32) bool {
	source := set.m.Load().(HashUint32Set)
	_, contains := source.data[item]
	return contains
}

func (set *HashSetAtomic) Remove(item string) {
	set.mu.Lock() // synchronize with other potential writers
	source := set.m.Load().(HashUint32Set)
	// hash item
	hashedKey := aeshash.Hash(item)
	delete(source.data, hashedKey)
	set.m.Store(source)
	set.mu.Unlock()
}

func (set *HashSetAtomic) LoadFromJSONArray(path string) {
	if path != "" {
		logger.Debug("loading hashset with json data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		var itemList []string
		unErr := json.Unmarshal(data, &itemList)
		if unErr != nil {
			logger.Error("could not unmarshal source data")
			return
		} else {
			set.LoadFromArray(itemList)
		}
	}
}

func (set *HashSetAtomic) LoadFromRaw(path string, splitChar string) {
	if path != "" {
		logger.Debug("loading hashset with raw data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		var itemList []string
		itemList = strings.Split(str.UnsafeString(data), splitChar)
		set.LoadFromArray(itemList)
	}
}

func (set *HashSetAtomic) LoadFromArray(data []string) {
	if data != nil {
		set.mu.Lock() // synchronize with other potential writers
		source := set.m.Load().(HashUint32Set)
		for _, v := range data {
			// hash item
			hashedKey := aeshash.Hash(v)
			source.data[hashedKey] = none
		}
		set.m.Store(source)
		set.mu.Unlock()
	}
}
