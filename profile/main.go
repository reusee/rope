package main

import (
	"bytes"
	"os"
	"runtime/pprof"

	"github.com/reusee/rope"
)

func main() {
	r := rope.NewFromBytes(bytes.Repeat([]byte("foobarbaz"), 1024))
	p, err := os.Create("pprof")
	if err != nil {
		panic(err)
	}
	pprof.StartCPUProfile(p)
	for i := 0; i < 10000000; i++ {
		r.Iter(0, func(bs []byte) bool {
			return true
		})
	}
	pprof.StopCPUProfile()
	p.Close()
}
