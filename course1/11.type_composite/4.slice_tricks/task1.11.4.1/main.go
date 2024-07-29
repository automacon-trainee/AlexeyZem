package main

func Cut(xs []int, start, end int) []int {
	if start > end || start < 0 || end > len(xs)-1 {
		return []int{}
	}
	return xs[start : end+1]
}

func main() {}
