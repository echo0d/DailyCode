package popcount

import "testing"

func BenchmarkSPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SPopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkCPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CPopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}
