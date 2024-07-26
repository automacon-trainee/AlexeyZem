package main

func Fibonacci(n int) int {
	if n == 1 {
		return 1
	}
	if n == 0 {
		return 0
	}
	cur := 1
	last := 0
	for i := 2; i <= n; i++ {
		cur, last = cur+last, cur
	}
	return cur
}

func main() {

}
