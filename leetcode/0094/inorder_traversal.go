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
