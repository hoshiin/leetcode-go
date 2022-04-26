package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func InorderTraversal(root *structures.TreeNode) []int {
	result := []int{}
	stack := []*structures.TreeNode{}
	node := root
	for node != nil || len(stack) > 0 {
		for node != nil {
			stack = append(stack, node)
			node = node.Left
		}
		node = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		node = node.Right
	}
	return result
}

func InorderTraversalRecursive(root *structures.TreeNode) []int {
	result := []int{}
	inorderInternal(root, &result)
	return result
}

func inorderInternal(root *structures.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		inorderInternal(root.Left, result)
	}
	*result = append(*result, root.Val)
	if root.Right != nil {
		inorderInternal(root.Right, result)
	}
}
