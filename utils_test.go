package rope

import "testing"

func TestEqual(t *testing.T) {
	r1 := &Rope{}
	var r2 *Rope
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		weight: 3,
	}
	r2 = &Rope{
		weight: 2,
	}
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		content: []byte("foo"),
	}
	r2 = &Rope{
		content: []byte("bar"),
	}
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		left: nil,
	}
	r2 = &Rope{
		left: &Rope{},
	}
	if r1.Equal(r2) {
		t.Fatal()
	}

	r1 = &Rope{
		right: new(Rope),
	}
	r2 = &Rope{
		right: &Rope{
			weight: 3,
		},
	}
	if r1.Equal(r2) {
		t.Fatal()
	}
}

func TestDump(t *testing.T) {
	r := &Rope{
		weight: 3,
		left: &Rope{
			weight:  3,
			content: []byte("foo"),
		},
		right: &Rope{
			weight:  4,
			content: []byte("barr"),
		},
	}
	r.Dump()
}
