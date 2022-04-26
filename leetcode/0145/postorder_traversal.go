package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func PostorderTraversal(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*structures.TreeNode{root}
	result := []int{}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, node.Val)
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
	}
	for i, j := 0, len(result)-1; i < len(result)/2; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}

func PostorderTraversalRecursive(root *structures.TreeNode) []int {
	result := []int{}
	postorderInternal(root, &result)
	return result
}

func postorderInternal(root *structures.TreeNode, result *[]int) {
	if root == nil {
		return
	}
	if root.Left != nil {
		postorderInternal(root.Left, result)
	}
	if root.Right != nil {
		postorderInternal(root.Right, result)
	}
	*result = append(*result, root.Val)
}
