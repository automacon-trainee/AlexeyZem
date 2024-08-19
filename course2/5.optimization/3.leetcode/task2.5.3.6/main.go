// Вам дан массив целых чисел nums без дубликатов. Максимальное бинарное
// дерево может быть построено рекурсивно из nums с помощью следующего
// алгоритма:
// Создайте корневой узел, значение которого является максимальным значением в nums.
// Рекурсивно постройте левое поддерево на подмассиве префикса левее максимального значения.
// Рекурсивно постройте правое поддерево на подмассиве суффикса правее
// максимального значения.
// Верните максимальное бинарное дерево, построенное из nums.
// Constraints:
// 1 <= nums.length <= 1000
// 0 <= nums[i] <= 1000
// All integers in nums are unique

package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	maxIndex := 0
	maxVal := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > maxVal {
			maxVal = nums[i]
			maxIndex = i
		}
	}
	left := constructMaximumBinaryTree(nums[:maxIndex])
	right := constructMaximumBinaryTree(nums[maxIndex+1:])
	return &TreeNode{Val: maxVal, Left: left, Right: right}
}

func main() {}
