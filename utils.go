package rope

import (
	"fmt"
	"strings"
)

var (
	p = fmt.Printf
)

func (r *Rope) Equal(r2 *Rope) bool {
	if r == nil && r2 == nil {
		return true
	}
	if r == nil && r2 != nil || r != nil && r2 == nil {
		return false
	}
	if r.Weight != r2.Weight {
		return false
	}
	if r.Text != r2.Text {
		return false
	}
	if !r.Left.Equal(r2.Left) {
		return false
	}
	if !r.Right.Equal(r2.Right) {
		return false
	}
	return true
}

func (r *Rope) Dump() {
	r.dump(0)
}

func (r *Rope) dump(level int) {
	p("%s%d |%s|\n", strings.Repeat("  ", level), r.Weight, r.Text)
	if r.Left != nil {
		r.Left.dump(level + 1)
	}
	if r.Right != nil {
		r.Right.dump(level + 1)
	}
}
