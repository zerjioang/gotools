// +build windows

package fastime

import "time"

func (t *FastTime) now() time.Time {
	wint := time.Now()
	t.nsec = wint.Nanosecond()
	t.sec = wint.Unix()
}
