package rope

import (
	"crypto/rand"
	"os"
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
