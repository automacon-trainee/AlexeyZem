package main

import (
	"fmt"
)

func Sum(nums ...int) int {
	res := 0
	for _, v := range nums {
		res += v
	}
	return res
}

func Mul(nums ...int) int {
	res := 1
	for _, v := range nums {
		res *= v
	}
	return res
}

func Sub(nums ...int) int {
	res := 0

	for i, v := range nums {
		if i == 0 {
			res = v
		} else {
			res -= v
		}
	}
	return res
}

func MathOperate(op func(nums ...int) int, nums ...int) int {
	return op(nums...)
}

func main() {
	fmt.Println(MathOperate(Sum, 1, 1, 3))
	fmt.Println(MathOperate(Sub, 1, 2, 3))
	fmt.Println(MathOperate(Mul, 1, 1, 3))
}
