// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package stringbank

import (
	"fmt"
	"runtime"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestStringbank(t *testing.T) {
	sb := Stringbank{}

	s1 := sb.Save("hello")
	s2 := sb.Save("goodbye")
	s3 := sb.Save("cheese")

	assert.Equal(t, "hello", sb.Get(s1))
	assert.Equal(t, "goodbye", sb.Get(s2))
	assert.Equal(t, "cheese", sb.Get(s3))
}

func TestStringbankSize(t *testing.T) {
	sb := Stringbank{}
	assert.Zero(t, sb.Size())
	sb.Save("hello")
	assert.Equal(t, stringbankSize, sb.Size())
}

func TestPackageBank(t *testing.T) {
	s1 := Save("hello")
	s2 := Save("goodbye")
	s3 := Save("cheese")

	assert.Equal(t, "hello", s1.String())
	assert.Equal(t, "goodbye", s2.String())
	assert.Equal(t, "cheese", s3.String())
}

func TestLengths(t *testing.T) {
	tests := []struct {
		len int
	}{
		{1},
		{127},
		{128},
		{254},
		{255},
		{256},
		{0xFFFFFFFFFF},
	}

	for _, test := range tests {
		t.Run(strconv.Itoa(test.len), func(t *testing.T) {
			buf := make([]byte, 10)

			l := writeLength(test.len, buf)
			assert.Equal(t, l, spaceForLength(test.len))
			len, lenlen := readLength(buf)
			assert.Equal(t, l, lenlen)
			assert.Equal(t, test.len, len)
		})
	}
}

func TestGC(t *testing.T) {
	sb := Stringbank{}
	for i := 0; i < 10000000; i++ {
		sb.Save(strconv.Itoa(i))
	}
	runtime.GC()

	start := time.Now()
	runtime.GC()
	assert.True(t, time.Since(start) < 1000*time.Microsecond)
	runtime.KeepAlive(sb)
}

func ExampleSave() {
	i := Save("hello")
	fmt.Println(i)
	// Output: hello
}

func ExampleStringbank() {
	sb := Stringbank{}
	i := sb.Save("goodbye")
	fmt.Println(sb.Get(i))
	// Output: goodbye
}
