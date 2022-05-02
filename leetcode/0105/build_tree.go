package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func BuildTree(preorder []int, inorder []int) *structures.TreeNode {
	if len(preorder) == 0 && len(inorder) == 0 {
		return nil
	}
	pivot := preorder[0]
	index := 0
	for inorder[index] != pivot {
		index++
	}
	left := BuildTree(preorder[1:index+1], inorder[:index])
	right := BuildTree(preorder[index+1:], inorder[index+1:])
	return &structures.TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}
