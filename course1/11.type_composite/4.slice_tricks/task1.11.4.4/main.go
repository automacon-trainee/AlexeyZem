package main

func RemoveIDX(x []int, idx int) []int {
	if idx < 0 || idx > len(x)-1 {
		return x
	}
	return append(x[:idx], x[idx+1:]...)
}

func main() {}
