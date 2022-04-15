package utils

import "testing"

func BenchmarkNewUUID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewUUID()
	}
}

func BenchmarkNewObjectID(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		NewOID()
	}
}
