package rope

import "testing"

func TestEqual(t *testing.T) {
	r1 := &Rope{}
	var r2 *Rope
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		Weight: 3,
	}
	r2 = &Rope{
		Weight: 2,
	}
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		Text: "foo",
	}
	r2 = &Rope{
		Text: "bar",
	}
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		Left: nil,
	}
	r2 = &Rope{
		Left: &Rope{},
	}
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		Right: new(Rope),
	}
	r2 = &Rope{
		Right: &Rope{
			Weight: 3,
		},
	}
	if r1.Equal(r2) {
		t.Fatal()
	}
}

func TestDump(t *testing.T) {
	r := &Rope{
		Weight: 3,
		Left: &Rope{
			Weight: 3,
			Text:   "foo",
		},
		Right: &Rope{
			Weight: 4,
			Text:   "barr",
		},
	}
	r.Dump()
}
