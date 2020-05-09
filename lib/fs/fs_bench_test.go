package fs_test

import (
	"bytes"
	"crypto/rand"
	"encoding/binary"
	"testing"

	"github.com/zerjioang/gotools/lib/fs"
)

func BenchmarkFs(b *testing.B) {
	b.Run("pagesize", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = fs.PageSize()
		}
	})
	b.Run("exists", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = fs.Exists("/tmp")
		}
	})
	b.Run("read-entropy", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = fs.ReadEntropy(rand.Reader, 16)
		}
	})
	b.Run("read-entropy-16", func(b *testing.B) {
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_, _ = fs.ReadEntropy16()
		}
	})
}

// binary writes bench test
func BenchmarkBinaryWrite(b *testing.B) {
	buf := &bytes.Buffer{}

	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf.Reset()

		for j := 0; j < 10; j++ {
			binary.Write(buf, binary.BigEndian, int32(j))
		}
	}
}

func BenchmarkBinaryPut(b *testing.B) {
	var writebuf [1024]byte

	b.ReportAllocs()
	b.SetBytes(1)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		buf := writebuf[0:0]

		for j := 0; j < 10; j++ {
			b := make([]byte, 4)
			binary.BigEndian.PutUint32(b, uint32(j))
			buf = append(buf, b...)
		}
	}
}
