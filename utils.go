package rope

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"log"
	"strings"
)

var (
	p = fmt.Printf
)

func (r *Rope) StructEqual(r2 *Rope) bool {
	if r == nil && r2 == nil {
		return true
	}
	if r == nil && r2 != nil || r != nil && r2 == nil {
		return false
	}
	if r.weight != r2.weight {
		return false
	}
	if !(bytes.Equal(r.content, r2.content)) {
		return false
	}
	if !r.left.StructEqual(r2.left) {
		return false
	}
	if !r.right.StructEqual(r2.right) {
		return false
	}
	return true
}

func (r *Rope) Dump() {
	r.dump(0, "")
}

func (r *Rope) dump(level int, prefix string) {
	p("%s%s%d |%s|\n", strings.Repeat("  ", level), prefix, r.weight, r.content)
	if r.left != nil {
		r.left.dump(level+1, "<")
	}
	if r.right != nil {
		r.right.dump(level+1, ">")
	}
}

func getRandomBytes(l int) []byte {
	bytes := make([]byte, l)
	n, err := rand.Read(bytes)
	if n != len(bytes) || err != nil {
		log.Fatalf("%d %v", n, err)
	}
	return bytes
}
