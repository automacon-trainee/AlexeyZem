// Дано корневое значение бинарного поискового дерева. Верните
// сбалансированное бинарное поисковое дерево с теми же значениями узлов. Если
// есть более одного ответа, верните любой из них
// Constraints:
// Количество узлов в дереве находится в диапазоне [1, 104].
// 1 <= Node.val <= 105

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := InOrder(root, []int{})
	return constructBalanceTree(nums)
}

func InOrder(root *TreeNode, sl []int) []int {
	if root == nil {
		return sl
	}
	if root.Left != nil {
		sl = InOrder(root.Left, sl)
	}
	sl = append(sl, root.Val)
	if root.Right != nil {
		sl = InOrder(root.Right, sl)
	}
	return sl
}

func constructBalanceTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	if len(nums) == 1 {
		return &TreeNode{Val: nums[0]}
	}
	left := constructBalanceTree(nums[:len(nums)/2])
	right := constructBalanceTree(nums[len(nums)/2+1:])
	return &TreeNode{nums[len(nums)/2], left, right}
}

func main() {}
