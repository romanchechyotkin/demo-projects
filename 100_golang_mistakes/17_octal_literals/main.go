package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	a := 0o10 // num started with 0 is an octal integer (base 8) => 010 in base 8 == 8 in base 10
	a = 010   // num started with 0 is an octal integer (base 8) => 010 in base 8 == 8 in base 10
	sum := 100 + a
	fmt.Println(sum) // 108

	// octal numbers are useful for opening file with permission
	file, err := os.OpenFile("foo", os.O_RDONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	_ = file

	// Using 0o as a prefix instead of only 0 means the same thing.
	file, err = os.OpenFile("foo", os.O_RDONLY, 0o644)
	if err != nil {
		log.Println(err)
	}
	_ = file

	binaryNum := 0b111
	fmt.Println(binaryNum) // 7

	hexadecimalNum := 0xF
	fmt.Println(hexadecimalNum) // 15

	imaginaryNum := 3i
	fmt.Println(imaginaryNum) // (0 + 3i)

	binaryNum = 0b0_1_010_1
	fmt.Println(binaryNum) // 21
}
