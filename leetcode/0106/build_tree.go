package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func BuildTree(inorder []int, postorder []int) *structures.TreeNode {
	if len(inorder) == 0 && len(postorder) == 0 {
		return nil
	}
	pivot := postorder[len(postorder)-1]
	pivotIndex := 0
	for inorder[pivotIndex] != pivot {
		pivotIndex++
	}

	left := BuildTree(inorder[:pivotIndex], postorder[:pivotIndex])
	right := BuildTree(inorder[pivotIndex+1:], postorder[pivotIndex:len(postorder)-1])
	root := &structures.TreeNode{Val: pivot, Left: left, Right: right}
	return root
}
