package main

import (
	"strings"
)

func CountWordOccurrences(text string) map[string]int {
	m := make(map[string]int)
	textSlice := strings.Split(text, " ")
	for _, word := range textSlice {
		m[word]++
	}
	return m
}

func main() {}
