// +build aix darwin dragonfly freebsd js,wasm linux nacl netbsd openbsd solaris

package fastime

import (
	"syscall"
)

func (t *FastTime) now() {
	t.nsec, t.sec = internalNow()
}

func internalNow() (uint32, int64) {
	var tv syscall.Timeval
	err := syscall.Gettimeofday(&tv)
	if err != nil {
	}
	return 0, syscall.TimevalToNsec(tv)
}
