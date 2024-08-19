// У вас есть n плиток, на каждой из которых напечатана одна буква tiles[i].
// Верните количество возможных непустых последовательностей букв, которые вы
// можете составить, используя буквы, напечатанные на этих плитках

package main

func backtrack(letters []int, res *int) {
	for i := 0; i < len(letters); i++ {
		if letters[i] > 0 {
			letters[i]--
			*res++
			backtrack(letters, res)
			letters[i]++
		}
	}
}

func numTilePossibilities(tiles string) int {
	res := 0
	countLetter := 26
	letters := make([]int, countLetter)
	for i := 0; i < len(tiles); i++ {
		letters[tiles[i]-'A']++
	}
	backtrack(letters, &res)
	return res
}

func main() {}
