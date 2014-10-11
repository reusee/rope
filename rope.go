package rope

import "bytes"

type Rope struct {
	weight  int
	left    *Rope
	right   *Rope
	content []byte
}

var MaxLengthPerNode = 128

func NewFromBytes(bs []byte) *Rope {
	if len(bs) == 0 {
		return nil
	}
	if len(bs) < MaxLengthPerNode {
		return &Rope{
			weight:  len(bs),
			content: bs,
		}
	}
	leftLen := len(bs) / 2
	return &Rope{
		weight: leftLen,
		left:   NewFromBytes(bs[:leftLen]),
		right:  NewFromBytes(bs[leftLen:]),
	}
}

func (r *Rope) Index(i int) byte {
	if i >= r.weight {
		return r.right.Index(i - r.weight)
	}
	if r.left != nil { // non leaf
		return r.left.Index(i)
	}
	// leaf
	return r.content[i]
}

func (r *Rope) Len() int {
	if r == nil {
		return 0
	}
	return r.weight + r.right.Len()
}

func (r *Rope) Bytes() []byte {
	buf := new(bytes.Buffer)
	r.collectBytes(buf)
	return buf.Bytes()
}

func (r *Rope) collectBytes(buf *bytes.Buffer) {
	if r == nil {
		return
	}
	if len(r.content) > 0 {
		buf.Write(r.content)
	} else {
		r.left.collectBytes(buf)
		r.right.collectBytes(buf)
	}
}

func (r *Rope) Concat(r2 *Rope) *Rope {
	return &Rope{
		weight: r.Len(),
		left:   r,
		right:  r2,
	}
}

func (r *Rope) Split(n int) (out1, out2 *Rope) {
	if len(r.content) > 0 { // leaf
		out1 = NewFromBytes(r.content[:n])
		out2 = NewFromBytes(r.content[n:])
	} else { // non leaf
		var r1 *Rope
		if n >= r.weight { // at right subtree
			r1, out2 = r.right.Split(n - r.weight)
			out1 = r.left.Concat(r1)
		} else { // at left subtree
			out1, r1 = r.left.Split(n)
			out2 = r1.Concat(r.right)
		}
	}
	return
}

func (r *Rope) Insert(n int, bs []byte) *Rope {
	r1, r2 := r.Split(n)
	return r1.Concat(NewFromBytes(bs)).Concat(r2)
}
