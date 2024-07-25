package main

import (
	"fmt"
	"math"
)

var (
	CalculateCircleArea    func(radius float64) float64
	CalculateRectangleArea func(width, height float64) float64
	CalculateTriangleArea  func(base, height float64) float64
)

func main() {
	CalculateCircleArea = func(radius float64) float64 {
		return math.Pi * radius * radius
	}
	CalculateRectangleArea = func(width, height float64) float64 {
		return width * height
	}
	CalculateTriangleArea = func(base, height float64) float64 {
		return base * height / 2
	}
	var a float64 = 10
	var b float64 = 20
	fmt.Println(CalculateCircleArea(a), CalculateRectangleArea(a, b), CalculateTriangleArea(a, b))
}
