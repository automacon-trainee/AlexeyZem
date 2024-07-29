package main

func Pop(xs []int) (first int, res []int) {
	if len(xs) == 0 {
		return 0, xs
	}
	first = xs[0]
	res = xs[1:]
	return first, res
}

func main() {}
