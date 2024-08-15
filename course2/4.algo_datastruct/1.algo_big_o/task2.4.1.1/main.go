package main

func FactorialRecursive(n int) int {
	if n < 0 {
		return 0
	}
	if n <= 1 {
		return 1
	}
	return n * FactorialRecursive(n-1)
}

func FactorialIterative(n int) int {
	if n < 0 {
		return 0
	}
	res := 1
	for i := 2; i <= n; i++ {
		res = res * i
	}
	return res
}

func main() {}
