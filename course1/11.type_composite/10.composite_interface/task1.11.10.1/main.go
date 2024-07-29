package main

import (
	"fmt"
)

// linter need any
// in this case interface{} == any
func getType(i any) string {
	switch i.(type) {
	case int:
		return "int"
	case string:
		return "string"
	case float64:
		return "float64"
	case bool:
		return "bool"
	case []int:
		return "[]int"
	case []float64:
		return "[]float64"
	default:
		return "Пустой интерфейс"
	}
}

func main() {
	fmt.Println(getType("int"))
	fmt.Println(getType(1))
	fmt.Println(getType([]int{1, 2, 3}))
	fmt.Println(getType(nil))
}
