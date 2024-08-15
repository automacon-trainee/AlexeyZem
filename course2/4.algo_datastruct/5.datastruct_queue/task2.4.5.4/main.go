package main

import (
	"fmt"
)

var stack []int

func Push(val int) {
	stack = append(stack, val)
}

func Pop() int {
	if len(stack) == 0 {
		return 0
	}
	val := stack[len(stack)-1]
	stack = stack[:len(stack)-1]
	return val
}

func main() {
	val1 := 10
	val2 := 20
	Push(val2)
	Push(val1)
	result := Pop() + Pop()
	Push(result)
	fmt.Println(stack[0])
}
