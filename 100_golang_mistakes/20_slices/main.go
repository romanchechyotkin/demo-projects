package main

import "fmt"

func main() {
	s := make([]int, 3, 6)
	fmt.Printf("length: %d; capacity: %d\n", len(s), cap(s))
	s = append(s, 1, 2, 3)
	fmt.Printf("length: %d; capacity: %d\n", len(s), cap(s))
	s = append(s, 7)
	fmt.Printf("length: %d; capacity: %d\n", len(s), cap(s))
}
