package main

func RemoveExtraMemory(xs []int) []int {
	if cap(xs) > len(xs) {
		xs = xs[:]
	}
	return xs
}

func main() {}
