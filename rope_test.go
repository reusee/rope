package rope

import (
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
