package utils

import "slices"

func NewStringSet(str ...string) map[string]struct{} {
	set := make(map[string]struct{})

	for _, txt := range str {
		set[txt] = struct{}{}
	}

	return set
}

func SortStringSet(set map[string]struct{}) []string {
	res := make([]string, 0, len(set))

	for k := range set {
		res = append(res, k)
	}

	slices.Sort(res)

	return res
}
