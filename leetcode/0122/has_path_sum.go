package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func HasPathSum(root *structures.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		return targetSum == 0
	}
	return HasPathSum(root.Left, targetSum) || HasPathSum(root.Right, targetSum)
}
func HasPathSumStack(root *structures.TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	stack := []*structures.TreeNode{root}
	sumStack := []int{targetSum - root.Val}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		sum := sumStack[len(sumStack)-1]
		sumStack = sumStack[:len(sumStack)-1]
		if node.Right == nil && node.Left == nil && sum == 0 {
			return true
		}
		if node.Right != nil {
			stack = append(stack, node.Right)
			sumStack = append(sumStack, sum-node.Right.Val)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
			sumStack = append(sumStack, sum-node.Left.Val)
		}
	}
	return false
}
