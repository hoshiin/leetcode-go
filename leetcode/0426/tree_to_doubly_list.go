package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func TreeToDoublyList(root *structures.TreeNode) *structures.TreeNode {
	if root == nil {
		return nil
	}
	var first, last *structures.TreeNode
	var dfs func(node *structures.TreeNode)

	dfs = func(node *structures.TreeNode) {
		if node != nil {
			dfs(node.Left)
			if last != nil {
				last.Right = node
				node.Left = last
			} else {
				first = node
			}
			last = node
			dfs(node.Right)
		}
	}
	dfs(root)
	last.Right = first
	first.Left = last
	return first
}
