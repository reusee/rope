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
	str := string(bytes)
	r := NewFromString(str)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		r.Index(3000)
	}
}
