package main

func InsertToStart(xs []int, nums ...int) []int {
	return append(nums, xs...)
}

func main() {}
