//1. Сумма самых глубоких листьев.
//Дан root дерево бинарных данных, верните сумму значений его самых глубоких листьев.
//Constraints:
//•
//Количество узлов в дереве находится в диапазоне [1, 104].
//•
//1 <= Node.val <= 100
//https://leetcode.com/problems/deepest-leaves-sum/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	cur := 0
	for len(queue) > 0 {
		cur = 0
		newLevel := []*TreeNode{}
		for _, node := range queue {
			if node.Left != nil {
				newLevel = append(newLevel, node.Left)
			}
			if node.Right != nil {
				newLevel = append(newLevel, node.Right)
			}
			cur += node.Val
		}
		queue = newLevel
	}
	return cur
}

func main() {}
