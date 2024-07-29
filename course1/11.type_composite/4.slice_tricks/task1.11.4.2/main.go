package main

func insertAfterIndex(s []int, index int, x ...int) []int {
	if index > len(s)-1 {
		return []int{}
	}

	x = append(x, s[index+1:]...)
	return append(s[:index+1], x...)
}

func main() {}
