package rope

import (
	"bytes"
	"testing"
)

func TestPersistent(t *testing.T) {
	bs := getRandomBytes(1024)

	// concat
	r := NewFromBytes(nil)
	var ropes []*Rope
	for i := 0; i < len(bs); i++ {
		r = r.Concat(NewFromBytes([]byte{bs[i]}))
		ropes = append(ropes, r)
	}
	for i := 0; i < len(bs); i++ {
		if !bytes.Equal(ropes[i].Bytes(), bs[:i+1]) {
			t.Fatal()
		}
	}

	// split
	r = NewFromBytes(bs)
	type ropeTuple struct {
		left, right *Rope
	}
	var ropeTuples []ropeTuple
	for i := 0; i < len(bs); i++ {
		r1, r2 := r.Split(i)
		ropeTuples = append(ropeTuples, ropeTuple{r1, r2})
	}
	for i := 0; i < len(bs); i++ {
		if !bytes.Equal(ropeTuples[i].left.Bytes(), bs[:i]) {
			t.Fatal()
		}
		if !bytes.Equal(ropeTuples[i].right.Bytes(), bs[i:]) {
			t.Fatal()
		}
	}

	// insert
	r = NewFromBytes(nil)
	ropes = ropes[0:0]
	for i := 0; i < len(bs); i++ {
		r = r.Insert(r.Len(), []byte{bs[i]})
		ropes = append(ropes, r)
	}
	for i := 0; i < len(bs); i++ {
		if !bytes.Equal(ropes[i].Bytes(), bs[:i+1]) {
			t.Fatal()
		}
	}

	// delete
	bs = []byte("foobarbaz")
	r = NewFromBytes(bs)
	ropes = ropes[0:0]
	for i := 0; i < len(bs); i++ {
		r = r.Delete(0, 1)
		ropes = append(ropes, r)
	}
	for i := 0; i < len(bs); i++ {
		if !bytes.Equal(ropes[i].Bytes(), bs[i+1:]) {
			t.Fatal()
		}
	}
}
