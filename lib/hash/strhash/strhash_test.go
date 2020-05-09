package strhash

import (
	"testing"
	"unsafe"
)

func TestStrHash(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		x1 := StrHash("foo-bar")
		t.Log(x1)
	})
	t.Run("complete-test", func(t *testing.T) {
		x0 := "abc"
		b := "b"
		b2 := "a" + b + "c"
		x1 := StrHash(x0)
		x2 := StrHash(x0)
		x3 := StrHash(b2)
		y0 := "def"
		y1 := StrHash(y0)
		y2 := StrHash(y0)

		if x1 != x2 || x2 != x3 {
			t.Errorf("x: should all be equal: %d - %d - %d", x1, x2, x3)
		}
		if !(unsafe.Pointer(&x0) != unsafe.Pointer(&b2) && x1 == x3) {
			t.Errorf("Hash should work on string value not pointer")
		}
		if y1 == x1 {
			t.Errorf("Different input strings should nearly 'always' have different hashes: %d - %d", x1, y1)
		}
		if y1 != y2 {
			t.Errorf("y: should all be equal: %d - %d", y1, y2)
		}
	})
}
