package main

import "log"

type Foo string
type Bar string

// TODO: benchmark with 3 ways to append to slice
// slice with 0 len
// slice with 0 len and given capacity
// slice with given len

func main() {
	foos := []Foo{"foo", "bar", "baz"}
	bars := convertWithCap(foos)
	log.Println(bars)
}

func convertWith0Len(foos []Foo) []Bar {
	bars := make([]Bar, 0)

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	return bars
}

func convertWithLen(foos []Foo) []Bar {
	bars := make([]Bar, len(foos))

	for i, foo := range foos {
		bars[i] = fooToBar(foo)
	}

	return bars
}

func convertWithCap(foos []Foo) []Bar {
	bars := make([]Bar, 0, len(foos))

	for _, foo := range foos {
		bars = append(bars, fooToBar(foo))
	}

	return bars
}

func fooToBar(foo Foo) Bar {
	return Bar(foo)
}
