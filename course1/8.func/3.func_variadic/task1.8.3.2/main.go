package main

import (
	"fmt"
	"strings"
)

func ConcatenateStrings(sep string, str ...string) string {
	res := "even: "
	for i := 0; i < len(str); i += 2 {
		res += str[i] + sep
	}
	res = strings.TrimSuffix(res, sep)
	res += ", odd: "
	for i := 1; i < len(str); i += 2 {
		res += str[i] + sep
	}
	res = strings.TrimSuffix(res, sep)
	return res
}

func main() {
	fmt.Println(ConcatenateStrings("-", "hello", "world", "how", "are", "you"))
}
