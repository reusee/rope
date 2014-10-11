package rope

type Rope struct {
	Weight int
	Left   *Rope
	Right  *Rope
	Bytes  []byte
}

var MaxLengthPerNode = 128

func NewFromBytes(bs []byte) *Rope {
	if len(bs) == 0 {
		return nil
	}
	if len(bs) < MaxLengthPerNode {
		return &Rope{
			Weight: len(bs),
			Bytes:  bs,
		}
	}
	leftLen := len(bs) / 2
	return &Rope{
		Weight: leftLen,
		Left:   NewFromBytes(bs[:leftLen]),
		Right:  NewFromBytes(bs[leftLen:]),
	}
}

func (r *Rope) Index(i int) byte {
	if i >= r.Weight {
		return r.Right.Index(i - r.Weight)
	}
	if r.Left != nil {
		return r.Left.Index(i)
	}
	return r.Bytes[i]
}

func (r *Rope) Len() int {
	if r == nil {
		return 0
	}
	return r.Weight + r.Right.Len()
}
