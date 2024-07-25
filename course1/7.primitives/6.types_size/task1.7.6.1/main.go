package main

import (
	"fmt"
	"unsafe"
)

func sizeOfBool(x bool) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfInt(x int) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfInt8(x int8) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfInt16(x int16) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfInt32(x int32) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfInt64(x int64) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfUint(x uint) int {
	return int(unsafe.Sizeof(x))
}

func sizeOfUint8(x uint8) int {
	return int(unsafe.Sizeof(x))
}

func main() {
	var (
		a bool
		b int
		c int8
		d int16
		e int32
		f int64
		g uint
		h uint8
	)
	fmt.Println("bool size:", sizeOfBool(a))
	fmt.Println("int size:", sizeOfInt(b))
	fmt.Println("int8 size", sizeOfInt8(c))
	fmt.Println("int16 size:", sizeOfInt16(d))
	fmt.Println("int32 size:", sizeOfInt32(e))
	fmt.Println("int64 size:", sizeOfInt64(f))
	fmt.Println("uint size:", sizeOfUint(g))
	fmt.Println("uint8 size:", sizeOfUint8(h))
}
