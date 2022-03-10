package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func DiameterOfBinaryTree(root *TreeNode) int {
	diameter := 0
	longestPath(root, &diameter)
	return diameter
}

func longestPath(node *TreeNode, result *int) int {
	if node == nil {
		return 0
	}
	leftPath := longestPath(node.Left, result)
	rightPath := longestPath(node.Right, result)

	*result = int(math.Max(float64(*result), float64(leftPath+rightPath)))

	return int(math.Max(float64(leftPath), float64(rightPath)) + 1)
}
