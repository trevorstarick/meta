package main

import (
	"os"

	"github.com/trevorstarick/meta"
)

func main() {
	for _, f := range os.Args[1:] {
		m, err := meta.Extract(f)
		if err != nil {
			panic(err)
		}

		_ = m

		// spew.Dump(m)
	}
}
