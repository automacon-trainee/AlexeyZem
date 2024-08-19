// В связанном списке размера n, где n является четным, i-й узел (с 0 индексом)
// связанного списка известен как близнец(twin) (n-1-i) узла, если 0 <= i <= (n / 2)
// Например, если n = 4, то узел 0 является близнецом узла 3, а узел 1 является
// близнецом узла 2. Это единственные узлы с близнецами для n = 4.
// Сумма близнецов(twin sum) определяется как сумма узла и его близнеца.
// Дан заголовок связанного списка с четной длиной, верните максимальную сумму
// близнецов связанного списка.
// Constraints:
// Число узлов в списке является четным целым числом в диапазоне [2, 105].
// 1 <= Node.val <= 105
//https://leetcode.com/problems/maximum-twin-sum-of-a-linked-list/

package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func pairSum(head *ListNode) int {
	sl := []int{}
	cur := head
	for cur != nil {
		sl = append(sl, cur.Val)
		cur = cur.Next
	}
	res := 0
	for i := 0; i < len(sl)/2; i++ {
		if sl[i]+sl[len(sl)-i-1] > res {
			res = sl[i] + sl[len(sl)-i-1]
		}
	}
	return res
}

func main() {}
