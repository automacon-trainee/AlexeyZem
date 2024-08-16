package main

func countRussianLetters(s string) map[rune]int {
	m := make(map[rune]int)
	for _, r := range s {
		if isRussianLetter(r) {
			m[r]++
		}
	}
	return m
}

func isRussianLetter(r rune) bool {
	return (r >= 'а' && r <= 'я') || (r >= 'А' && r <= 'Я')
}

func main() {}
