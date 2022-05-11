package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func SearchBST(root *structures.TreeNode, val int) *structures.TreeNode {
	if root == nil {
		return nil
	}
	if root.Val == val {
		return root
	}
	left := SearchBST(root.Left, val)
	if left != nil {
		return left
	}
	right := SearchBST(root.Right, val)
	return right
}
