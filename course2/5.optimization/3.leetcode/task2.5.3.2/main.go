// Дана класс с m студентами и n экзаменами. Вам дана матрица целых чисел 0
// индексированных m x n, где каждая строка представляет одного студента, а score [i][j] обозначает балл,
// который i- тый студент получил на j - том экзамене. Матрица
// оценок содержит только уникальные целые числа.
// Вам также дано целое число k. Отсортируйте студентов (т. e. Строки матрицы) по их
// баллам в k - ом (0- индексированном) экзамене от высшего к низшему.
// Верните матрицу после сортировк
// Constraints:
// m == score.length
// n == score[i].length
// 1 <= m, n <= 250
// 1 <= score[i][j] <= 105
// score consists of distinct integers.
// 0 <= k < n
// https://leetcode.com/problems/sort-the-students-by-their-kth-score/
package main

import (
	"sort"
)

func sortTheStudents(score [][]int, k int) [][]int {
	sort.Slice(score, func(i, j int) bool {
		return score[i][k] < score[j][k]
	})
	return score
}

func main() {}
