package main

import (
	"fmt"
	"log"
)

func main() {
	var num int
	var err error
	fmt.Println(num) // 0

	var condition bool

	if condition {
		n, err := getNumber2(condition)
		if err != nil {
			log.Println(err)
		}
		num = n
	} else {
		num, err = getNumber2(condition)
	}
	if err != nil {
		log.Println(err)
	}

	fmt.Println(num) // 0
}

func getNumber2(cond bool) (int, error) {
	if cond {
		return 1, nil
	}
	return 2, nil
}
