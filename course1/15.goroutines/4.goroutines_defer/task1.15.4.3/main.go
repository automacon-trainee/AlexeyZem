package main

import (
	"fmt"
)

func main() {
	ch := make(chan string)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			close(ch)
		}
	}()
	myPanic(ch)
	fmt.Println(<-ch)
}

func myPanic(_ chan string) {
	panic("my panic message")
}
