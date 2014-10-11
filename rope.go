package rope

import (
	"bytes"
	"math"
)

type Rope struct {
	height  int
	weight  int
	left    *Rope
	right   *Rope
	content []byte
}

var MaxLengthPerNode = 128

func NewFromBytes(bs []byte) (ret *Rope) {
	if len(bs) == 0 {
		return nil
	}
	slots := make([]*Rope, 32)
	for blockIndex := 0; blockIndex < len(bs)/MaxLengthPerNode; blockIndex++ {
		block := bs[blockIndex*MaxLengthPerNode : (blockIndex+1)*MaxLengthPerNode]
		r := &Rope{
			height:  1,
			weight:  MaxLengthPerNode,
			content: block,
		}
		slotIndex := 0
		for slots[slotIndex] != nil {
			r = &Rope{
				height: slotIndex + 2,
				weight: (1 << uint(slotIndex)) * MaxLengthPerNode,
				left:   slots[slotIndex],
				right:  r,
			}
			slots[slotIndex] = nil
			slotIndex++
		}
		slots[slotIndex] = r
	}
	tailStart := len(bs) / MaxLengthPerNode * MaxLengthPerNode
	if tailStart < len(bs) {
		ret = &Rope{
			height:  1,
			weight:  len(bs) - tailStart,
			content: bs[tailStart:],
		}
	}
	for _, c := range slots {
		if c != nil {
			if ret == nil {
				ret = c
			} else {
				ret = c.Concat(ret)
			}
		}
	}
	return
}

//TODO func NewFromReader
//TODO func (r *Rope) Read

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

func (r *Rope) Concat(r2 *Rope) (ret *Rope) {
	ret = &Rope{
		height: 0,
		weight: r.Len(),
		left:   r,
		right:  r2,
	}
	if ret.left != nil {
		ret.height = ret.left.height
	}
	if ret.right != nil && ret.right.height > ret.height {
		ret.height = ret.right.height
	}
	ret.height++
	// check and rebalance
	l := int((math.Ceil(math.Log2(float64((ret.Len()/MaxLengthPerNode)+1))) + 1) * 1.5)
	if ret.height > l {
		ret = NewFromBytes(ret.Bytes())
	}
	return
}

func (r *Rope) Split(n int) (out1, out2 *Rope) {
	if r == nil {
		return
	}
	if len(r.content) > 0 { // leaf
		if n > len(r.content) { // offset overflow
			n = len(r.content)
		}
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

func (r *Rope) Delete(n, l int) *Rope {
	r1, r2 := r.Split(n)
	_, r2 = r2.Split(l)
	return r1.Concat(r2)
}

func (r *Rope) Sub(n, l int) []byte {
	buf := new(bytes.Buffer)
	r.sub(n, l, buf)
	return buf.Bytes()
}

func (r *Rope) sub(n, l int, buf *bytes.Buffer) {
	if len(r.content) > 0 { // leaf
		end := n + l
		if end > len(r.content) {
			end = len(r.content)
		}
		buf.Write(r.content[n:end])
	} else { // non leaf
		if n >= r.weight { // start at right subtree
			r.right.sub(n-r.weight, l, buf)
		} else { // start at left subtree
			r.left.sub(n, l, buf)
			if n+l > r.weight {
				r.right.sub(0, n+l-r.weight, buf)
			}
		}
	}
}
