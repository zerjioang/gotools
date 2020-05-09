package rng

import (
	"crypto/rand"
	"runtime"
)

func init() {
	var b [256]byte
	var count int64
	cur := count & 0xF
	count++
	if cur == 0 {
		rand.Read(b[:])
	}
	start, end := cur*16, (cur+1)*16
	_ = b[start:end]
	runtime.Gosched()
}
