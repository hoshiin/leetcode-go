package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func CountUnivalSubtrees(root *structures.TreeNode) int {
	count := 0
	if root == nil {
		return 0
	}
	isUni(root, &count)
	return count
}

func isUni(node *structures.TreeNode, count *int) bool {
	if node.Left == nil && node.Right == nil {
		*count++
		return true
	}
	isUnival := true
	if node.Left != nil {
		isUnival = isUni(node.Left, count) && node.Left.Val == node.Val
	}
	if node.Right != nil {
		isUnival = isUni(node.Right, count) && isUnival && node.Right.Val == node.Val
	}
	if !isUnival {
		return false
	}
	*count++
	return true
}

func CountUnivalSubtreesOptimized(root *structures.TreeNode) int {
	if root == nil {
		return 0
	}
	_, count := isValidSubtree(root, root.Val, 0)
	return count
}

func isValidSubtree(node *structures.TreeNode, val int, count int) (bool, int) {
	if node == nil {
		return true, count
	}

	isValidLeft, count := isValidSubtree(node.Left, node.Val, count)
	isValidRight, count := isValidSubtree(node.Right, node.Val, count)
	if !isValidLeft || !isValidRight {
		return false, count
	}

	count++
	return node.Val == val, count
}
