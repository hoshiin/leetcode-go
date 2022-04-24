package leetcode

import (
	"github.com/hoshiin/leetcode-go/structures"
)

func PreorderTraversal(root *structures.TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := []*structures.TreeNode{}
	stack = append(stack, root)
	result := []int{}
	for len(stack) != 0 {
		visit := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		result = append(result, visit.Val)
		if visit.Right != nil {
			stack = append(stack, visit.Right)
		}
		if visit.Left != nil {
			stack = append(stack, visit.Left)
		}
	}
	return result
}
