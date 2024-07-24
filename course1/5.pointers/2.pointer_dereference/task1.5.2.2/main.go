package main

import "fmt"

func Factorial(n *int) int {
	if n == nil {
		return 0
	}
	res := 1
	for i := *n; i > 1; i-- {
		res *= i
	}
	return res
}

func isPalindrome(str *string) bool {
	for i := 0; i < len(*str)/2; i++ {
		if (*str)[i] != (*str)[len(*str)-i-1] {
			return false
		}
	}
	return true
}

func CountOccurrences(numbers *[]int, target *int) int {
	result := 0
	for i := 0; i < len(*numbers); i++ {
		if (*numbers)[i] == *target {
			result++
		}
	}
	return result
}

func revereString(str *string) string {
	sl := []byte(*str)
	for i := 0; i < len(sl)/2; i++ {
		sl[i], sl[len(sl)-1-i] = sl[len(sl)-1-i], sl[i]
	}
	return string(sl)
}
func main() {
	a := 5
	fmt.Println(Factorial(&a))

	str := "abba"
	fmt.Println(isPalindrome(&str))

	fmt.Println(revereString(&str))

	fmt.Println(CountOccurrences(&[]int{1, 2, 3, 3, 2, 1, 2, 3, 4, 5, 5, 6, 5, 5, 10}, &a))
}
