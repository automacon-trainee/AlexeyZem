package main

func Factorial(n int) int {
	if n == 0 {
		return 1
	}

	if n < 3 {
		return n
	}

	return n * Factorial(n-1)
}

func main() {

}
