package main

import (
	"fmt"
	"unsafe"
)

func binaryStringToFloat(s string) float32 {
	var number uint32
	for _, char := range s {
		number <<= 1
		if char == '1' {
			number += 1
		}
		if char != '1' && char != '0' {
			return 0
		}
	}
	floatNumber := *(*float32)(unsafe.Pointer(&number))
	return floatNumber
}

func main() {
	s := "00111110001000000000000000000000"
	fmt.Println(binaryStringToFloat(s))
}
