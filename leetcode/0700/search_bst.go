package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func SearchBST(root *structures.TreeNode, val int) *structures.TreeNode {
	if root == nil || root.Val == val {
		return root
	}
	if val < root.Val {
		return SearchBST(root.Left, val)
	}
	return SearchBST(root.Right, val)
}
