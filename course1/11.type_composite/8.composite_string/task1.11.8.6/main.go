package main

import (
	"strings"
)

var filter = map[rune]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
	'y': {},
	'а': {},
	'о': {},
	'и': {},
	'у': {},
	'е': {},
	'ы': {},
	'э': {},
	'я': {},
	'ю': {},
}

func CountVowels(s string) int {
	count := 0
	str := strings.ToLower(s)
	for _, val := range str {
		if _, ok := filter[val]; ok {
			count++
		}
	}
	return count
}

func main() {}
