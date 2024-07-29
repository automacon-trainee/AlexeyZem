package main

import (
	"strings"
)

func filterSentence(sentence string, filter map[string]bool) string {
	res := ""
	text := strings.Split(sentence, " ")
	for _, word := range text {
		if _, ok := filter[word]; !ok {
			res += word + " "
		}
	}
	return strings.TrimSpace(res)
}

func main() {}
