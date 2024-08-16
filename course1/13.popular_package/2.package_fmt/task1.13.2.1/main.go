package main

import (
	"fmt"
)

func generateMathStrings(operands []int, operator string) string {
	res := ""
	if len(operands) == 0 {
		return res
	}
	mathResult := operands[0]
	for i := 0; i < len(operands); i++ {
		if i != len(operands)-1 {
			res += fmt.Sprintf("%v %s ", operands[i], operator)
		} else {
			res += fmt.Sprintf("%v = ", operands[i])
		}
		if i != 0 {
			switch operator {
			case "+":
				mathResult += operands[i]
			case "-":
				mathResult -= operands[i]
			case "*":
				mathResult *= operands[i]
			case "/":
				mathResult /= operands[i]
			default:
				return ""
			}
		}
	}
	return res + fmt.Sprintf("%v", mathResult)
}

func main() {}
