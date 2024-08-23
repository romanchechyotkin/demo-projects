package main

import (
	"fmt"
)

func main() {
	var n float32 = 1.0001
	fmt.Println(n * n) // 1.0002 instead of 1.00020001
}
