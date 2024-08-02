package main

import (
	"fmt"
	"os"
)

func writeToFile(file *os.File, data string) error {
	defer file.Close()
	_, err := file.WriteString(data)
	return err
}

func main() {
	err := writeToFile(os.Stdout, "hello world\n")
	if err != nil {
		fmt.Println("writeToFile err:", err)
	}
}
