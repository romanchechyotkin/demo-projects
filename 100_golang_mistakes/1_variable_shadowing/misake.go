package main

import (
	"fmt"
	"log"
)

func main() {
	var num int
	fmt.Println(num) // 0

	var condition bool

	if condition {
		num, err := getNumber1(condition)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(num) // 1

	} else {
		num, err := getNumber1(condition)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(num) // 2
	}
	fmt.Println(num) // 0
}

func getNumber1(cond bool) (int, error) {
	if cond {
		return 1, nil
	}
	return 2, nil
}
