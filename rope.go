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
	var slotIndex int
	var r *Rope
	for blockIndex := 0; blockIndex < len(bs)/MaxLengthPerNode; blockIndex++ {
		r = &Rope{
			height:  1,
			weight:  MaxLengthPerNode,
			content: bs[blockIndex*MaxLengthPerNode : (blockIndex+1)*MaxLengthPerNode],
		}
		slotIndex = 0
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
	r.Iter(0, func(bs []byte) bool {
		buf.Write(bs)
		return true
	})
	return buf.Bytes()
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
	r.Iter(n, func(bs []byte) bool {
		if l >= len(bs) {
			buf.Write(bs)
			l -= len(bs)
			return true
		} else {
			buf.Write(bs[:l])
			return false
		}
	})
	return buf.Bytes()
}

func (r *Rope) Iter(offset int, fn func([]byte) bool) {
	r.iter(offset, fn)
}

func (r *Rope) iter(offset int, fn func([]byte) bool) bool {
	if r == nil {
		return true
	}
	if len(r.content) > 0 { // leaf
		if offset < len(r.content) {
			if !fn(r.content[offset:]) {
				return false
			}
		}
	} else { // non leaf
		if offset >= r.weight { // start at right subtree
			if !r.right.iter(offset-r.weight, fn) {
				return false
			}
		} else { // start at left subtree
			if !r.left.iter(offset, fn) {
				return false
			}
			if !r.right.iter(0, fn) {
				return false
			}
		}
	}
	return true
}
