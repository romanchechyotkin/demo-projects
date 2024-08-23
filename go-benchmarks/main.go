package main

import (
	"log"

	"benchmarks/slices"
)

func main() {
	var in = []string{"a", "b", "c", "d", "e", "f"}
	out := slices.CopySlice(in)
	log.Println(out)
}
