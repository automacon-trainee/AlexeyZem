package main

import (
	"strings"
)

func CountWordsInText(txt string, words []string) map[string]int {
	counts := make(map[string]int)
	txt = strings.ToLower(txt)
	for _, w := range words {
		counts[strings.ToLower(w)] = 0
	}
	sl := strings.Fields(txt)
	for _, w := range sl {
		if _, ok := counts[w]; ok {
			counts[w]++
		}
	}
	return counts
}

func main() {}
