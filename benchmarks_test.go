package rope

import (
	"crypto/rand"
	"log"
	"testing"
)

const benchBytesLen = 409600

func getBenchBytes() []byte {
	bytes := make([]byte, benchBytesLen)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		log.Fatalf("%d %v", n, err)
	}
	return bytes
}

func getBenchRope() *Rope {
	return NewFromBytes(getBenchBytes())
}

func BenchmarkNew(b *testing.B) {
	bytes := getBenchBytes()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.SetBytes(benchBytesLen)
		NewFromBytes(bytes)
	}
}

func BenchmarkIndex(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Index(300000)
	}
}

func BenchmarkLen(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Len()
	}
}

func BenchmarkBytes(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Bytes()
	}
}

func BenchmarkConcat(b *testing.B) {
	r1 := getBenchRope()
	r2 := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r1.Concat(r2)
	}
}

func BenchmarkSplit(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Split(300000)
	}
}

func BenchmarkInsert(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Insert(300000, []byte("foobar"))
	}
}

func BenchmarkDelete(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Delete(300000, 400000)
	}
}

func BenchmarkSub(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Sub(300000, 128)
	}
}

func BenchmarkNewRuneReader(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.NewRuneReader()
	}
}

func BenchmarkIter(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Iter(func(*Rope) {})
	}
}
