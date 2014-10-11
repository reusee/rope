package rope

import (
	"regexp"
	"testing"
	"unicode/utf8"
)

func TestRuneReader(t *testing.T) {
	r := NewFromBytes([]byte(`我能吞zuo下da玻si璃而不伤身体`))
	reader := r.NewRuneReader()
	type info struct {
		r rune
		n int
	}
	var res []info
	for {
		c, n, err := reader.ReadRune()
		if err != nil {
			break
		}
		res = append(res, info{c, n})
	}
	expected := []info{
		{'我', 3},
		{'能', 3},
		{'吞', 3},
		{'z', 1},
		{'u', 1},
		{'o', 1},
		{'下', 3},
		{'d', 1},
		{'a', 1},
		{'玻', 3},
		{'s', 1},
		{'i', 1},
		{'璃', 3},
		{'而', 3},
		{'不', 3},
		{'伤', 3},
		{'身', 3},
		{'体', 3},
	}
	if len(res) != len(expected) {
		p("%v\n%v\n", res, expected)
		t.Fatal()
	}
	for i, o := range res {
		if o.r != expected[i].r || o.n != expected[i].n {
			t.Fatal()
		}
	}

	r = NewFromBytes(nil)
	reader = r.NewRuneReader()
	c, n, err := reader.ReadRune()
	if c != utf8.RuneError || n != 1 || err == nil {
		t.Fatal()
	}
}

func TestRuneRegexp(t *testing.T) {
	r := NewFromBytes([]byte(`我能吞zuo下da玻si璃而不伤身体`))
	reader := r.NewRuneReader()
	loc := regexp.MustCompile(`[a-z]+`).FindReaderIndex(reader)
	if string(r.Sub(loc[0], loc[1]-loc[0])) != "zuo" {
		t.Fatal()
	}
}
