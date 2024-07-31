package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func GenerateRandomString(length int) string {
	str := strings.Builder{}
	var maxGenerate int64 = 100

	for i := 0; i < length; i++ {
		num, _ := rand.Int(rand.Reader, big.NewInt(maxGenerate))
		str.WriteRune(rune(num.Int64()))
	}
	return str.String()
}

func main() {}
