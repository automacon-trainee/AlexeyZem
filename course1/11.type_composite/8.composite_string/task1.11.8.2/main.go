package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func getStringsHeader(str string) reflect.StringHeader {
	return *(*reflect.StringHeader)(unsafe.Pointer(&str))
}

func main() {
	s := "Hello, World!"
	header := getStringsHeader(s)
	fmt.Println(header.Data)
	fmt.Println(header.Len)

}
