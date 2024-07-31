package main

import (
	"regexp"
)

func isValidEmail(email string) bool {
	re := regexp.MustCompile(`([a-zA-Z0-9._-]+@[a-zA-Z0-9._-]+\.[a-zA-Z]+)`)
	return re.MatchString(email)
}

func main() {}
