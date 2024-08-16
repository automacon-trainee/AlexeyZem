package main

import (
	"crypto/rand"
	"math/big"
	"strings"
)

func generateActivationKey() string {
	count := 4
	sl := make([]string, count)
	var maxNum int64 = 'z' - '0'
	for i := 0; i < count; i++ {
		str := strings.Builder{}
		for j := 0; j < count; j++ {
			run := ']'
			for (run > '9' && run < 'A') || (run > 'Z' && run < 'a') {
				num, _ := rand.Int(rand.Reader, big.NewInt(maxNum))
				run = '0' + rune(num.Int64())
			}
			str.WriteRune(run)
		}
		sl[i] = str.String()
	}
	return strings.Join(sl, "-")
}

func main() {}
