package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1[1:2]
	s3 := append(s2, 10)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}
