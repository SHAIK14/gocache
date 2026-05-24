package store

import "testing"

func BenchmarkSet(b *testing.B) {
	s := NewStore()
	for i := 0; i < b.N; i++ {
		s.Set("key", "value")
	}
}

func BenchmarkGet(b *testing.B) {
	s := NewStore()
	s.Set("key", "value")
	for i := 0; i < b.N; i++ {
		s.Get("key")
	}
}
