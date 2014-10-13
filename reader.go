package rope

import (
	"errors"
	"unicode/utf8"
)

type RuneReader struct {
	slices [][]byte
	cur    []byte
}

func (r *RuneReader) ReadRune() (rune, int, error) {
	if len(r.cur) < 4 && len(r.slices) > 0 {
		r.cur = append(r.cur, r.slices[0]...)
		r.slices = r.slices[1:]
	}
	c, n := utf8.DecodeRune(r.cur)
	if c == utf8.RuneError {
		return utf8.RuneError, 1, errors.New("decode error")
	}
	r.cur = r.cur[n:]
	return c, n, nil
}

func (r *Rope) NewRuneReader() *RuneReader {
	reader := new(RuneReader)
	r.Iter(func(leaf *Rope) bool {
		reader.slices = append(reader.slices, leaf.content)
		return true
	})
	return reader
}
