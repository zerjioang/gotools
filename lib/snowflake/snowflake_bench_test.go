package snowflake

import "testing"

func BenchmarkParseBase32(b *testing.B) {

	node, _ := NewNode(1)
	sf := node.Generate()
	b32i := sf.Base32()

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = ParseBase32([]byte(b32i))
	}
}
func BenchmarkBase32(b *testing.B) {

	node, _ := NewNode(1)
	sf := node.Generate()

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sf.Base32()
	}
}
func BenchmarkParseBase58(b *testing.B) {

	node, _ := NewNode(1)
	sf := node.Generate()
	b58 := sf.Base58()

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = ParseBase58([]byte(b58))
	}
}
func BenchmarkBase58(b *testing.B) {

	node, _ := NewNode(1)
	sf := node.Generate()

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sf.Base58()
	}
}
func BenchmarkGenerate(b *testing.B) {

	node, _ := NewNode(1)

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = node.Generate()
	}
}

func BenchmarkGenerateMaxSequence(b *testing.B) {

	NodeBits = 1
	StepBits = 21
	node, _ := NewNode(1)

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = node.Generate()
	}
}

func BenchmarkUnmarshal(b *testing.B) {
	// Generate the ID to unmarshal
	node, _ := NewNode(1)
	id := node.Generate()
	bytes, _ := id.MarshalJSON()

	var id2 ID

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = id2.UnmarshalJSON(bytes)
	}
}

func BenchmarkMarshal(b *testing.B) {
	// Generate the ID to marshal
	node, _ := NewNode(1)
	id := node.Generate()

	b.ReportAllocs()
	b.SetBytes(1)

	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_, _ = id.MarshalJSON()
	}
}
