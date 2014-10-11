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

func (r *Rope) putRuneReader(reader *RuneReader) {
	if r == nil {
		return
	}
	if len(r.content) > 0 {
		reader.slices = append(reader.slices, r.content)
	} else {
		r.left.putRuneReader(reader)
		r.right.putRuneReader(reader)
	}
}

func (r *Rope) NewRuneReader() *RuneReader {
	reader := new(RuneReader)
	r.putRuneReader(reader)
	return reader
}
