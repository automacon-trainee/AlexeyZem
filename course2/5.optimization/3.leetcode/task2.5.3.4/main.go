// Дано root звено бинарного поискового дерева (BST), преобразуйте его в большее
// дерево так, чтобы каждый ключ исходного BST был изменен на исходный ключ
// плюс сумма всех ключей, больших исходного ключа в BST.
// В качестве напоминания, бинарное поисковое дерево
// это дерево, которое удовлетворяет следующим ограничениям:
// Левое поддерево узла содержит только узлы с ключами меньше ключа узла.
// Правое поддерево узла содержит только узлы с ключами больше ключа узла.
// Оба поддерева должны быть также бинарными поисковыми деревьями.
// Constraints:
// Диапазон значений узлов находится в диапазоне [1, 100]
// 0 <= Node.val <= 100
// Все значения в дереве уникальны.
// https://leetcode.com/problems/binary-search-tree-to-greater-sum-tree/

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstToGst(root *TreeNode) *TreeNode {
	bstToGstNode(root, 0)
	return root
}
func bstToGstNode(root *TreeNode, sum int) int {
	if root == nil {
		return sum
	}
	right := bstToGstNode(root.Right, sum)
	root.Val += right
	left := bstToGstNode(root.Left, root.Val)
	return left
}

func main() {}
