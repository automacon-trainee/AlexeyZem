package main

func ReplaceSymbols(s string, old, newR rune) string {
	sl := []rune(s)
	for i, v := range sl {
		if v == old {
			sl[i] = newR
		}
	}
	return string(sl)
}

func main() {}
