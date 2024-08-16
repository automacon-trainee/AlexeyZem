package main

func countUniqueUTF8Characters(s string) int {
	m := make(map[rune]struct{})
	for _, v := range s {
		if _, ok := m[v]; !ok {
			m[v] = struct{}{}
		}
	}

	return len(m)
}

func main() {}
