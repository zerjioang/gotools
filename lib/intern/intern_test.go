// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package intern_test

import (
	"reflect"
	"strconv"
	"testing"
	"unsafe"

	"github.com/zerjioang/gotools/lib/intern"

	"github.com/stretchr/testify/assert"
)

func TestBasic(t *testing.T) {
	i := intern.New(16)

	hat := i.Deduplicate("hat")
	sat := i.Deduplicate("sat")
	hat2 := i.Deduplicate("hat")

	if hat != hat2 || hat != "hat" {
		t.Errorf("Hat is wrong. Have %s and %s", hat, hat2)
	}

	if sat != "sat" {
		t.Errorf("sat is wrong. Have %s", sat)
	}

	if datapointer(hat) != datapointer(hat2) {
		t.Errorf("hat pointers differ")
	}
}

func TestGrowth(t *testing.T) {
	in := intern.New(15)

	for i := 0; i < 1000; i++ {
		val := strconv.Itoa(i)
		assert.Equal(t, val, in.Deduplicate(val))
	}

	for i := 0; i < 1000; i++ {
		val := strconv.Itoa(i)
		assert.Equal(t, val, in.Deduplicate(val))
	}
}

func TestGrowth2(t *testing.T) {
	in := intern.New(15)

	for i := 0; i < 1000; i++ {
		val := strconv.Itoa(i)
		assert.Equal(t, val, in.Deduplicate(val))
		assert.Equal(t, val, in.Deduplicate(val))
	}
}

func TestNoNew(t *testing.T) {
	in := &intern.Intern{}

	for i := 0; i < 1000; i++ {
		val := strconv.Itoa(i)
		assert.Equal(t, val, in.Deduplicate(val))
		assert.Equal(t, val, in.Deduplicate(val))
	}
}

func datapointer(val string) uintptr {
	return (*reflect.StringHeader)(unsafe.Pointer(&val)).Data
}

func TestResize(t *testing.T) {
	i := intern.New(16)

	for k := 0; k < 2; k++ {
		for j := 0; j < 256; j++ {
			i.Deduplicate(strconv.Itoa(j))
		}
		if i.Len() != 256 {
			t.Errorf("expected 256 unique strings. Have %d", i.Len())
		}
		if i.Cap() != 512 {
			t.Errorf("expected 512 capacity. Have %d", i.Cap())
		}
	}
}
