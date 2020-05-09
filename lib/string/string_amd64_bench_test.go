package string

import (
	"strings"
	"testing"

	"github.com/zerjioang/gotools/lib/logger"
)

func BenchmarkAssembly(b *testing.B) {
	b.Run("is-digit-go", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = isDigitGo('0')
		}
	})
	b.Run("is-digit-asm", func(b *testing.B) {
		logger.Enabled(false)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsDigit('0')
		}
	})
	b.Run("is-digit-array-go", func(b *testing.B) {
		logger.Enabled(false)
		example := []byte("1485545485")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = isDigitArrayGo(example)
		}
	})
	b.Run("is-digit-array-asm", func(b *testing.B) {
		logger.Enabled(false)
		example := []byte("1485545485")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			_ = IsNumericArray(example)
		}
	})
	b.Run("lowercase-go", func(b *testing.B) {
		logger.Enabled(false)
		example := "CONVERT TO LOWER CASE"
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			strings.ToLower(example)
		}
	})
	b.Run("lowercase-go-custom", func(b *testing.B) {
		logger.Enabled(false)
		example := []byte("CONVERT TO LOWER CASE")
		s := NewWith(example)
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			s.LowerCase()
		}
	})
	b.Run("lowercase-asm", func(b *testing.B) {
		logger.Enabled(false)
		example := []byte("CONVERT TO LOWER CASE")
		b.ReportAllocs()
		b.SetBytes(1)
		b.ResetTimer()
		for n := 0; n < b.N; n++ {
			LowerCase(example)
		}
	})
}

func isDigitGo(b byte) bool {
	return b >= '0' && b <= '9'
}

func isDigitArrayGo(data []byte) bool {
	for idx := 0; idx < len(data); idx++ {
		b := data[idx]
		if !(b >= '0' && b <= '9') {
			return false
		}
	}
	return true
}
