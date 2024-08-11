package main

import "fmt"

func Add(a, b int) *int {
	result := new(int)
	*result = a + b
	return result
}
func Max(numbers []int) *int {
	if len(numbers) == 0 {
		return nil
	}
	j := 0
	for i, n := range numbers {
		if n > numbers[j] {
			j = i
		}
	}
	return &numbers[j]
}
func IsPrime(number int) *bool {
	result := new(bool)

	if number <= 0 {
		*result = false
		return result
	}

	for i := 2; i < number; i++ {
		if number%i == 0 {
			*result = false
			return result
		}
	}

	*result = true
	return result
}
func ConcatenateStrings(strs []string) *string {
	result := new(string)
	for _, str := range strs {
		*result += str
	}
	return result
}
func main() {
	fmt.Println(*Max([]int{3, 2, 3, 5, 7, 11, 13, 10, 100}))
	fmt.Println(*Add(1, 2))
	fmt.Println(*ConcatenateStrings([]string{"jfgef", "wjf", "hifb"}))
	fmt.Println(*IsPrime(10))
}
