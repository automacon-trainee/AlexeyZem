// Вам дан заголовок связанного списка, который содержит ряд целых чисел,
// разделенных нулями. Начало и конец связного списка будет иметь Node.val == 0.
// Для каждых двух последовательных нулей объединяйте все узлы, лежащие между
// ними в один узел, чье значение является суммой всех объеденных узлов.
// Измененный список не должен содержать никаких нулей.
// Верните заголовок измененного связного списка
// Constraints:
// Количество узлов в списке находится в диапазоне [3, 2 * 105].
// 0 <= Node.val <= 1000
// Нет двух последовательных узлов с Node.val == 0.
// Начало и конец связного списка имеют Node.val == 0.
// https://leetcode.com/problems/merge-nodes-in-between-zeros/

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeNodes(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Next != nil && cur.Next.Val != 0 {
			cur.Val += cur.Next.Val
			cur.Next = cur.Next.Next
		} else if cur.Next != nil && cur.Next.Next != nil {
			cur = cur.Next
		} else if cur.Next != nil && cur.Next.Next == nil {
			cur.Next = cur.Next.Next
		}
	}
	return head
}

func main() {
	n1 := &ListNode{Val: 0}
	n2 := &ListNode{Val: 3}
	n3 := &ListNode{Val: 1}
	n4 := &ListNode{Val: 0}
	n5 := &ListNode{Val: 4}
	n6 := &ListNode{Val: 5}
	n7 := &ListNode{Val: 2}
	n8 := &ListNode{Val: 0}
	n1.Next = n2
	n2.Next = n3
	n3.Next = n4
	n4.Next = n5
	n5.Next = n6
	n6.Next = n7
	n7.Next = n8
	h := mergeNodes(n1)
	fmt.Println(h.Val)
}
