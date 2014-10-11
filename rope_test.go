package rope

import (
	"bytes"
	"crypto/rand"
	mrand "math/rand"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	MaxLengthPerNode = 8
	os.Exit(m.Run())
}

func TestNewFromBytes(t *testing.T) {
	// nil bytes
	r := NewFromBytes([]byte{})
	if r != nil {
		t.Fatal()
	}

	// short bytes
	r = NewFromBytes([]byte(`foo`))
	if !r.StructEqual(&Rope{
		weight:  3,
		content: []byte("foo"),
	}) {
		r.Dump()
		t.Fatal()
	}

	// long bytes
	r = NewFromBytes([]byte(`foobarbaz`))
	if !r.StructEqual(&Rope{
		weight: 4,
		left: &Rope{
			weight:  4,
			content: []byte("foob"),
		},
		right: &Rope{
			weight:  5,
			content: []byte("arbaz"),
		},
	}) {
		r.Dump()
		t.Fatal()
	}
}

func TestIndex(t *testing.T) {
	bs := []byte(`abcdefghijklmnopqrstuvwxyz0123456789`)
	r := NewFromBytes(bs)
	for i := 0; i < len(bs); i++ {
		if r.Index(i) != bs[i] {
			t.Fatal()
		}
	}

	bs = make([]byte, 4096)
	n, err := rand.Read(bs)
	if n != len(bs) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	r = NewFromBytes(bs)
	for i := 0; i < len(bs); i++ {
		if r.Index(i) != bs[i] {
			t.Fatal()
		}
	}
}

func TestLen(t *testing.T) {
	if NewFromBytes([]byte{}).Len() != 0 {
		t.Fatal()
	}
	for i := 0; i < 1024; i++ {
		n := mrand.Intn(2048)
		if NewFromBytes(bytes.Repeat([]byte("x"), n)).Len() != n {
			t.Fatal()
		}
	}
}

func TestBytes(t *testing.T) {
	for i := 0; i < 1024; i++ {
		bs := make([]byte, i)
		n, err := rand.Read(bs)
		if n != len(bs) || err != nil {
			t.Fatalf("%d %v", n, err)
		}
		r := NewFromBytes(bs)
		if !bytes.Equal(r.Bytes(), bs) {
			t.Fatal()
		}
	}
}

func TestConcat(t *testing.T) {
	for i := 0; i < 1024; i++ {
		bs1 := make([]byte, i)
		n, err := rand.Read(bs1)
		if n != len(bs1) || err != nil {
			t.Fatalf("%d %v", n, err)
		}
		bs2 := make([]byte, i)
		n, err = rand.Read(bs2)
		if n != len(bs2) || err != nil {
			t.Fatalf("%d %v", n, err)
		}
		r1 := NewFromBytes(bs1)
		r2 := NewFromBytes(bs2)
		if !bytes.Equal(r1.Concat(r2).Bytes(), bytes.Join([][]byte{bs1, bs2}, nil)) {
			t.Fatal()
		}
	}
}

func TestSplit(t *testing.T) {
	r := NewFromBytes([]byte(`foobarbaz`))
	r1, r2 := r.Split(0)
	if !bytes.Equal(r1.Bytes(), []byte{}) {
		t.Fatal()
	}
	if !bytes.Equal(r2.Bytes(), []byte("foobarbaz")) {
		t.Fatal()
	}
	r1, r2 = r.Split(9)
	if !bytes.Equal(r1.Bytes(), []byte("foobarbaz")) {
		t.Fatal()
	}
	if !bytes.Equal(r2.Bytes(), []byte{}) {
		t.Fatal()
	}

	bs := make([]byte, 2048)
	n, err := rand.Read(bs)
	if n != len(bs) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	r = NewFromBytes(bs)
	for i := 0; i <= len(bs); i++ {
		r1, r2 := r.Split(i)
		if !bytes.Equal(r1.Bytes(), bs[:i]) {
			t.Fatal()
		}
		if !bytes.Equal(r2.Bytes(), bs[i:]) {
			t.Fatal()
		}
	}
}

func TestInsert(t *testing.T) {
	r := NewFromBytes([]byte(`foobar`))
	if string(r.Insert(0, []byte(`baz`)).Bytes()) != "bazfoobar" {
		t.Fatal()
	}
	if string(r.Insert(6, []byte(`baz`)).Bytes()) != "foobarbaz" {
		t.Fatal()
	}
	if string(r.Insert(2, []byte(`baz`)).Bytes()) != "fobazobar" {
		t.Fatal()
	}

	bs := make([]byte, 2048)
	n, err := rand.Read(bs)
	if n != len(bs) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	r = NewFromBytes(bs)
	for i := 0; i <= len(bs); i++ {
		bs1 := r.Insert(i, []byte("FOOBARBAZ")).Bytes()
		bs2 := bytes.Join([][]byte{bs[:i], []byte("FOOBARBAZ"), bs[i:]}, nil)
		if !bytes.Equal(bs1, bs2) {
			p("%s %s\n", bs1, bs2)
			t.Fatal()
		}
	}
}
