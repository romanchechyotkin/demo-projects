package main

import "fmt"

func main() {
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println(dst) // [], cause len(dst) == 0

	dst2 := make([]int, len(src))
	copy(dst2, src)
	fmt.Println(dst2) // [1, 2, 3]

}
