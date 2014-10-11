package rope

import (
	"crypto/rand"
	mrand "math/rand"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	MaxLengthPerNode = 8
	os.Exit(m.Run())
}

func TestNewFromString(t *testing.T) {
	// nil string
	r := NewFromString("")
	if r != nil {
		t.Fatal()
	}

	// short string
	r = NewFromString(`foo`)
	if !r.Equal(&Rope{
		Weight: 3,
		Text:   "foo",
	}) {
		r.Dump()
		t.Fatal()
	}

	// long string
	r = NewFromString(`foobarbaz`)
	if !r.Equal(&Rope{
		Weight: 4,
		Left: &Rope{
			Weight: 4,
			Text:   "foob",
		},
		Right: &Rope{
			Weight: 5,
			Text:   "arbaz",
		},
	}) {
		r.Dump()
		t.Fatal()
	}
}

func TestIndex(t *testing.T) {
	str := `abcdefghijklmnopqrstuvwxyz0123456789`
	r := NewFromString(str)
	for i := 0; i < len(str); i++ {
		if r.Index(i) != str[i] {
			t.Fatal()
		}
	}

	bytes := make([]byte, 4096)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	str = string(bytes)
	r = NewFromString(str)
	for i := 0; i < len(str); i++ {
		if r.Index(i) != str[i] {
			t.Fatal()
		}
	}
}

func TestLen(t *testing.T) {
	if NewFromString("").Len() != 0 {
		t.Fatal()
	}
	for i := 0; i < 1024; i++ {
		n := mrand.Intn(2048)
		if NewFromString(strings.Repeat("x", n)).Len() != n {
			t.Fatal()
		}
	}
}
