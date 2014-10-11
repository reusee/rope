package rope

import (
	"bytes"
	"crypto/rand"
	mrand "math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	MaxLengthPerNode = 8
	os.Exit(m.Run())
}

func TestNewFromBytes(t *testing.T) {
	// nil bytes
	r := NewFromBytes([]byte{})
	if r != nil {
		t.Fatal()
	}

	// short bytes
	r = NewFromBytes([]byte(`foo`))
	if !r.Equal(&Rope{
		weight:  3,
		content: []byte("foo"),
	}) {
		r.Dump()
		t.Fatal()
	}

	// long bytes
	r = NewFromBytes([]byte(`foobarbaz`))
	if !r.Equal(&Rope{
		weight: 4,
		left: &Rope{
			weight:  4,
			content: []byte("foob"),
		},
		right: &Rope{
			weight:  5,
			content: []byte("arbaz"),
		},
	}) {
		r.Dump()
		t.Fatal()
	}
}

func TestIndex(t *testing.T) {
	bs := []byte(`abcdefghijklmnopqrstuvwxyz0123456789`)
	r := NewFromBytes(bs)
	for i := 0; i < len(bs); i++ {
		if r.Index(i) != bs[i] {
			t.Fatal()
		}
	}

	bytes := make([]byte, 4096)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	r = NewFromBytes(bytes)
	for i := 0; i < len(bytes); i++ {
		if r.Index(i) != bytes[i] {
			t.Fatal()
		}
	}
}

func TestLen(t *testing.T) {
	if NewFromBytes([]byte{}).Len() != 0 {
		t.Fatal()
	}
	for i := 0; i < 1024; i++ {
		n := mrand.Intn(2048)
		if NewFromBytes(bytes.Repeat([]byte("x"), n)).Len() != n {
			t.Fatal()
		}
	}
}
