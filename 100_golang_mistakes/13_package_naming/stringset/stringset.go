package stringset

import "slices"

type Set map[string]struct{}

func New(str ...string) Set {
	set := make(map[string]struct{})

	for _, txt := range str {
		set[txt] = struct{}{}
	}

	return set
}

func Sort(set Set) []string {
	res := make([]string, 0, len(set))

	for k := range set {
		res = append(res, k)
	}

	slices.Sort(res)

	return res
}
