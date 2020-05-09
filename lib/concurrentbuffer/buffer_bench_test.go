// Copyright gotools (https://github.com/zerjioang/gotools)
// SPDX-License-Identifier: GNU GPL v3

package concurrentbuffer

import (
	"bytes"
	"testing"
)

var (
	testMessage = []byte("hello")
)

func BenchmarkConcurrentBuffer(b *testing.B) {
	b.Run("instantiate", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = NewConcurrentBuffer()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = NewConcurrentBufferPtr()
			}
		})
	})

	b.Run("read", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.Read(testMessage)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.Read(testMessage)
			}
		})
	})

	b.Run("write", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.Write(testMessage)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.Write(testMessage)
			}
		})
	})

	b.Run("string", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.String()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.String()
			}
		})
	})

	b.Run("string", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Bytes()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Bytes()
			}
		})
	})

	b.Run("cap", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Cap()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Cap()
			}
		})
	})

	b.Run("grow", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Grow(1)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Grow(1)
			}
		})
	})

	b.Run("len", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Len()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Len()
			}
		})
	})

	b.Run("next", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Next(0)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_ = buf.Next(0)
			}
		})
	})

	b.Run("read-byte", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadByte()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadByte()
			}
		})
	})

	b.Run("read-bytes", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadBytes(byte(0))
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadBytes(byte(0))
			}
		})
	})

	b.Run("read-from", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			reader := bytes.NewReader(testMessage)
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadFrom(reader)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			reader := bytes.NewReader(testMessage)
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadFrom(reader)
			}
		})
	})

	b.Run("read-rune", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _, _ = buf.ReadRune()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _, _ = buf.ReadRune()
			}
		})
	})

	b.Run("read-string", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadString(byte(0))
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				_, _ = buf.ReadString(byte(0))
			}
		})
	})

	b.Run("reset", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Reset()
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Reset()
			}
		})
	})

	b.Run("truncate", func(b *testing.B) {
		b.Run("struct", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBuffer()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Truncate(0)
			}
		})
		b.Run("ptr", func(b *testing.B) {
			b.ReportAllocs()
			b.SetBytes(1)
			buf := NewConcurrentBufferPtr()
			_, _ = buf.Write(testMessage)
			b.ResetTimer()
			for n := 0; n < b.N; n++ {
				buf.Truncate(0)
			}
		})
	})
}
