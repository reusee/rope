package rope

import (
	"crypto/rand"
	"testing"
)

func BenchmarkIndex(b *testing.B) {
	bytes := make([]byte, 4096)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		b.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bytes)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Index(3000)
	}
}
