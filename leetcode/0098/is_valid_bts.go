package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func IsValidBST(root *structures.TreeNode) bool {
	return validate(root, nil, nil)
}

func validate(root *structures.TreeNode, low, high *int) bool {
	if root == nil {
		return true
	}
	if (low != nil && root.Val <= *low) || (high != nil && root.Val >= *high) {
		return false
	}
	return validate(root.Right, &root.Val, high) && validate(root.Left, low, &root.Val)
}
