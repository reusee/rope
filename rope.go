package rope

type Rope struct {
	Weight int
	Left   *Rope
	Right  *Rope
	Text   string
}

var MaxLengthPerNode = 128

func NewFromString(str string) *Rope {
	if len(str) == 0 {
		return nil
	}
	if len(str) < MaxLengthPerNode {
		return &Rope{
			Weight: len(str),
			Text:   str,
		}
	}
	leftLen := len(str) / 2
	return &Rope{
		Weight: leftLen,
		Left:   NewFromString(str[:leftLen]),
		Right:  NewFromString(str[leftLen:]),
	}
}

func (r *Rope) Index(i int) uint8 {
	if i >= r.Weight {
		return r.Right.Index(i - r.Weight)
	}
	if r.Left != nil {
		return r.Left.Index(i)
	}
	return r.Text[i]
}
