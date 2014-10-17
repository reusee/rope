package rope

import (
	"bytes"
	"crypto/rand"
	"log"
	"math"
	mrand "math/rand"
	"os"
	"testing"
)

func getRandomBytes(l int) []byte {
	bytes := make([]byte, l)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		log.Fatalf("%d %v", n, err)
	}
	return bytes
}

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
		weight: 8,
		left: &Rope{
			weight:  8,
			content: []byte("foobarba"),
		},
		right: &Rope{
			weight:  1,
			content: []byte("z"),
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

func TestDelete(t *testing.T) {
	r := NewFromBytes([]byte(`foobarbaz`))
	cases := []struct {
		start, length int
		str           string
	}{
		{0, 0, "foobarbaz"},
		{0, 1, "oobarbaz"},
		{0, 2, "obarbaz"},
		{0, 9, ""},
		{1, 1, "fobarbaz"},
		{1, 2, "fbarbaz"},
		{4, 4, "foobz"},
		{5, 3, "foobaz"},
		{9, 0, "foobarbaz"},
	}
	for _, c := range cases {
		s := string(r.Delete(c.start, c.length).Bytes())
		if s != c.str {
			p("%s %s\n", s, c.str)
			t.Fatal()
		}
	}

	bs := make([]byte, 128)
	n, err := rand.Read(bs)
	if n != len(bs) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	r = NewFromBytes(bs)
	for i := 0; i <= len(bs); i++ {
		for j := 0; j <= len(bs); j++ {
			k := i + j
			if k > len(bs) {
				k = len(bs)
			}
			expected := bytes.Join([][]byte{bs[:i], bs[k:]}, nil)
			bs1 := r.Delete(i, j).Bytes()
			if !bytes.Equal(bs1, expected) {
				t.Fatal()
			}
		}
	}
}

func TestSub(t *testing.T) {
	bs := make([]byte, 128)
	n, err := rand.Read(bs)
	if n != len(bs) || err != nil {
		t.Fatalf("%d %v", n, err)
	}
	r := NewFromBytes(bs)
	for i := 0; i < len(bs); i++ {
		for j := 0; j < len(bs); j++ {
			end := i + j
			if end > len(bs) {
				end = len(bs)
			}
			expected := bs[i:end]
			bs1 := r.Sub(i, j)
			if !bytes.Equal(bs1, expected) {
				t.Fatal()
			}
		}
	}
}

func TestBalance(t *testing.T) {
	r := NewFromBytes(nil)
	n := 4096
	for i := 0; i < n; i++ {
		r = r.Concat(NewFromBytes([]byte("x")))
	}
	if r.Len() != n {
		t.Fatal()
	}
	maxHeight := int(math.Log2(float64(n/MaxLengthPerNode))+1) * 2
	if r.height > maxHeight {
		t.Fatal()
	}
}

func TestIter(t *testing.T) {
	r := NewFromBytes(bytes.Repeat([]byte("foobarbaz"), 512))
	r.Iter(0, func([]byte) bool {
		return true
	})

	n := 0
	r.Iter(0, func([]byte) bool {
		n++
		if n == 3 {
			return false
		}
		return true
	})
	if n != 3 {
		t.Fatal()
	}

	for i := 0; i < r.Len(); i++ {
		n := 0
		r.Iter(i, func(bs []byte) bool {
			n += len(bs)
			return true
		})
		if n != r.Len()-i {
			t.Fatal()
		}
	}

	r.Iter(r.weight, func([]byte) bool {
		return false
	})
}

func TestIterNodes(t *testing.T) {
	r := NewFromBytes(bytes.Repeat([]byte("foobarbaz"), 512))
	buf := new(bytes.Buffer)
	r.iterNodes(func(n *Rope) bool {
		if len(n.content) > 0 {
			buf.Write(n.content)
		}
		return true
	})
	if !bytes.Equal(buf.Bytes(), r.Bytes()) {
		t.Fatal()
	}
}

func TestIterBackward(t *testing.T) {
	r := NewFromBytes(nil)
	r.IterBackward(0, func([]byte) bool {
		t.Fatal()
		return true
	})

	bs := bytes.Repeat([]byte("foobarbaz"), 512)
	r = NewFromBytes(bs)
	for i := 0; i <= r.Len(); i++ {
		buf := new(bytes.Buffer)
		r.IterBackward(i, func(bs []byte) bool {
			buf.Write(bs)
			return true
		})
		if !bytes.Equal(buf.Bytes(), reversedBytes(bs[:i])) {
			t.Fatal()
		}
	}

	n := 0
	r.IterBackward(r.Len(), func([]byte) bool {
		n++
		return false
	})
	if n != 1 {
		t.Fatal()
	}

	n = 0
	r.IterBackward(r.Len(), func(bs []byte) bool {
		n += len(bs)
		if n > r.Len()-r.weight {
			return false
		}
		return true
	})
	if n != 520 {
		t.Fatal()
	}

	n = 0
	r.IterBackward(r.weight-8, func(bs []byte) bool {
		n++
		return false
	})
	if n != 1 {
		t.Fatal()
	}
}

func TestIterRune(t *testing.T) {
	r := NewFromBytes([]byte("我能吞zuo下da玻si璃而不伤身体"))
	expected := []rune{
		'我', '能', '吞', 'z', 'u', 'o',
		'下', 'd', 'a', '玻', 's', 'i', '璃',
		'而', '不', '伤', '身', '体',
	}
	i := 0
	r.IterRune(0, func(c rune, l int) bool {
		if c != expected[i] {
			t.Fatal()
		}
		i++
		return true
	})
	if i != len(expected) {
		t.Fatal()
	}

	expected = []rune{
		'z', 'u', 'o', '下', 'd', 'a', '玻',
		's', 'i', '璃', '而', '不', '伤', '身', '体',
	}
	i = 0
	r.IterRune(9, func(c rune, l int) bool {
		if c != expected[i] {
			t.Fatal()
		}
		i++
		return true
	})
	if i != len(expected) {
		t.Fatal()
	}

	expected = []rune{
		'z', 'u', 'o', '下', 'd', 'a', '玻',
		's', 'i', '璃',
	}
	i = 0
	r.IterRune(9, func(c rune, l int) bool {
		p("%d\n", i)
		if c != expected[i] {
			t.Fatal()
		}
		i++
		if c == '璃' {
			return false
		}
		return true
	})
	if i != len(expected) {
		t.Fatal()
	}

	r.IterRune(1, func(c rune, l int) bool {
		t.Fatal()
		return true
	})
}
