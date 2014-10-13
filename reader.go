package rope

import (
	"errors"
	"unicode/utf8"
)

type RuneReader struct {
	bs       chan []byte
	sigClose chan struct{}
	cur      []byte
}

func (r *RuneReader) ReadRune() (rune, int, error) {
	if len(r.cur) < 4 {
		bs, ok := <-r.bs
		if ok {
			r.cur = append(r.cur, bs...)
		}
	}
	c, n := utf8.DecodeRune(r.cur)
	if c == utf8.RuneError {
		return utf8.RuneError, 1, errors.New("decode error")
	}
	r.cur = r.cur[n:]
	return c, n, nil
}

func (r *RuneReader) Close() {
	close(r.sigClose)
}

func (r *Rope) NewRuneReader() *RuneReader {
	reader := &RuneReader{
		bs:       make(chan []byte),
		sigClose: make(chan struct{}),
	}
	go func() {
		r.Iter(0, func(bs []byte) bool {
			select {
			case reader.bs <- bs:
				return true
			case <-reader.sigClose:
				return false
			}
		})
		close(reader.bs)
	}()
	return reader
}
