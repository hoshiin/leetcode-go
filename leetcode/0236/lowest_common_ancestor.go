package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func LowestCommonAncestor(root, p, q *structures.TreeNode) *structures.TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := LowestCommonAncestor(root.Left, p, q)
	right := LowestCommonAncestor(root.Right, p, q)
	if left != nil && right != nil {
		return root
	}
	if left == nil {
		return right
	}
	return left
}
