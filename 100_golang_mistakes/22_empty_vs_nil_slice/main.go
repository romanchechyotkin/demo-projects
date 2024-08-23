package main

import "fmt"

func main() {
	var sl []string
	log(1, sl)

	sl = []string(nil)
	log(2, sl)

	sl = []string{}
	log(3, sl)

	sl = make([]string, 0)
	log(4, sl)

	ints := getSlice()
	if len(ints) != 0 { // check length, not slice != nil
		toDoSmth()
	}

}

func toDoSmth() {}

func log(i int, s []string) {
	fmt.Printf("%d: empty=%t\tnil=%t\n", i, len(s) == 0, s == nil)
}

func getSlice() []int {
	sl := make([]int, 0)

	if true {
		return nil
	}

	return sl
}
