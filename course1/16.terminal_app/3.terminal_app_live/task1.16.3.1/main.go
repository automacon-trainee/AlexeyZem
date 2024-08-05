package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Print("\033[H\033[2J")
		hour, minute, second := time.Now().Clock()
		year, month, day := time.Now().Date()
		fmt.Printf("Текущее время: %v:%v:%v\n", hour, minute, second)
		fmt.Printf("Текущая дата: %v-%d-%v\n", year, month, day)
		time.Sleep(1 * time.Second)
	}
}
