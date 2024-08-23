package slices

func CopySlice(in []string) []string {
	out := make([]string, len(in))
	for idx, val := range in {
		out[idx] = val
	}
	return out
}
