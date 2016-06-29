package rope

import (
	"bytes"
	"testing"
)

const benchBytesLen = 1 * 1024 * 1024

func getBenchBytes() []byte {
	return getRandomBytes(benchBytesLen)
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
		r.Index(512 * 1024)
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
		r.Split(512 * 1024)
	}
}

func BenchmarkInsert(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Insert(128*1024, []byte("foobar"))
	}
}

func BenchmarkDelete(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Delete(128*1024, 200*1024)
	}
}

func BenchmarkSub(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Sub(128*1024, 1024)
	}
}

func BenchmarkIter(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.Iter(0, func([]byte) bool {
			return true
		})
	}
}

func BenchmarkRebalance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := NewFromBytes(nil)
		for j := 0; j < 2048; j++ {
			r = r.Concat(NewFromBytes([]byte{'x'}))
		}
	}
}

func BenchmarkIterBackward(b *testing.B) {
	r := getBenchRope()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.IterBackward(r.Len(), func([]byte) bool {
			return true
		})
	}
}

func BenchmarkIterRune(b *testing.B) {
	r := NewFromBytes(bytes.Repeat([]byte("foobarbaz"), 1024))
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.IterRune(0, func(ru rune, l int) bool {
			return true
		})
	}
}
