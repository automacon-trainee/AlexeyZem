package main

func Shift(xs []int) (first int, res []int) {
	if len(xs) == 0 {
		return 0, []int{}
	}
	res = append([]int{xs[len(xs)-1]}, xs[:len(xs)-1]...)
	first = res[0]
	return first, res
}

func main() {}
