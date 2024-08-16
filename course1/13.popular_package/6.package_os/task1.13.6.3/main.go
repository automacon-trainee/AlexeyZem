package main

import (
	"fmt"
	"os"
)

func ReadString(filepath string) string {
	length := 1024
	res := ""
	buf := make([]byte, length)
	file, err := os.Open(filepath)
	if err != nil {
		return ""
	}
	defer file.Close()
	for n, err := file.Read(buf); err == nil; n, err = file.Read(buf) {
		res += string(buf[:n])
	}
	return res
}

func main() {
	fmt.Println(ReadString("aa.txt"))
}
