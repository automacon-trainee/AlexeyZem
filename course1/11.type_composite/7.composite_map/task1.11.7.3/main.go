package main

import (
	"strings"
)

func createUniqueText(text string) string {
	m := make(map[string]int)
	res := ""
	textSlice := strings.Split(text, " ")
	for _, word := range textSlice {
		m[word]++
		if m[word] == 1 {
			res += word + " "
		}
	}
	return strings.TrimSpace(res)
}

func main() {}
