// Дан ориентированный ациклический граф, с n вершинами, пронумерованными от
// 0 до n-1, и массив ребер, где edges [i] = [fromi, toi] представляет собой
// направленное ребро от узла fromi к узлу toi .
// Найдите самый маленький набор вершин, из которых достижимы все узлы в графе.
// Гарантируется существование уникального решения.
// Constraints:
// 2 <= n <= 10^5
// 1 <= edges.length <= min(10^5, n * (n-1) / 2)
// edges[i].length == 2
// 0 <= fromi, toi < n
// All pairs (fromi, toi) are distinct.
// https://leetcode.com/problems/minimum-number-of-vertices-to-reach-all-nodes/

package main

func findSmallestSetOfVertices(n int, edges [][]int) []int {
	res := make([]int, 0)
	from, to := make(map[int]struct{}, n), make(map[int]struct{}, n)
	for _, edge := range edges {
		from[edge[0]] = struct{}{}
		to[edge[1]] = struct{}{}
	}
	for ver := range from {
		if _, ok := to[ver]; !ok {
			res = append(res, ver)
		}
	}
	return res
}

func main() {}
