// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package hashset

import (
	"encoding/json"
	"io/ioutil"
	"strings"

	"github.com/zerjioang/gotools/util/str"

	"github.com/zerjioang/gotools/lib/logger"
)

// HashSet WORM (Write Once Read Many)
// is an unsafe data structure in where data is loaded once
// for example, at init() functions an it is readed, after loading many times
// avoiding further modifications after being loaded.
type HashSetWORM struct {
	set HashSet
}

func NewHashSetWORM() HashSetWORM {
	hs := HashSetWORM{}
	hs.set = NewHashSet()
	return hs
}

func NewHashSetWORMPtr() *HashSetWORM {
	hs := NewHashSetWORM()
	return &hs
}

func (s *HashSetWORM) Add(item string) {
	s.set.data[item] = none
}

func (s *HashSetWORM) Clear() {
	s.set = NewHashSet()
}

func (s *HashSetWORM) ClearFast() {
	s.set.ClearFast()
}

func (s HashSetWORM) Contains(item string) bool {
	_, found := s.set.data[item]
	return found
}

func (s HashSetWORM) MatchAny(item string) bool {
	found := false
	for k := range s.set.data {
		found = strings.Contains(item, k)
		if found {
			break
		}
	}
	return found
}

func (s *HashSetWORM) Remove(item string) {
	delete(s.set.data, item)
}

func (s *HashSetWORM) Size() int {
	l := len(s.set.data)
	return l
}

func (s *HashSetWORM) LoadFromJSONArray(path string) {
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
			s.LoadFromArray(itemList)
		}
	}
}

func (s *HashSetWORM) LoadFromRaw(path string, splitChar string) {
	if path != "" {
		logger.Debug("loading hashset with raw data")
		data, err := ioutil.ReadFile(path)
		if err != nil {
			logger.Error("could not read source data")
			return
		}
		var itemList []string
		itemList = strings.Split(str.UnsafeString(data), splitChar)
		s.LoadFromArray(itemList)
	}
}

// content loaded via this method will make to allocate data on the heap
func (s *HashSetWORM) LoadFromArray(data []string) {
	if data != nil {
		for _, v := range data {
			s.set.data[v] = none
		}
	}
}
