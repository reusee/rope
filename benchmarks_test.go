package rope

import (
	"crypto/rand"
	"testing"
)

func BenchmarkIndex(b *testing.B) {
	bytes := make([]byte, 409600)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		b.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Index(300000)
	}
}

func BenchmarkSplit(b *testing.B) {
	bytes := make([]byte, 409600)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		b.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Split(300000)
	}
}

func BenchmarkInsert(b *testing.B) {
	bytes := make([]byte, 409600)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		b.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Insert(300000, []byte("foobar"))
	}
}

func BenchmarkDelete(b *testing.B) {
	bytes := make([]byte, 409600)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		b.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Delete(300000, 400000)
	}
}

func BenchmarkSub(b *testing.B) {
	bytes := make([]byte, 409600)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		b.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Sub(300000, 128)
	}
}
