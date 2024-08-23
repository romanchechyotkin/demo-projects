package main

import (
	"fmt"
	"math"
)

func main() {
	var num int64 = math.MaxInt64
	fmt.Println(num) // 9223372036854775807
	num++
	fmt.Println(num) // -9223372036854775808

	var n int32 = math.MaxInt32
	newN := Inc32(n)
	fmt.Println(newN)
}

func Inc32(n int32) int32 {
	if n == math.MaxInt32 {
		panic("int32 overflow")
	}
	return n + 1
}

func AddInt(a, b int) int {
	if a > math.MaxInt-b {
		panic("int overflow")
	}

	return a + b
}

func MultiplicationInt(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}

	res := a * b
	if a == 1 || b == 1 {
		return res
	}

	if a == math.MinInt || b == math.MinInt {
		panic("int overflow")
	}

	if res/b != a {
		panic("integer overflow")
	}

	return res
}
