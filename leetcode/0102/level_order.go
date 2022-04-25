package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func LevelOrder(root *structures.TreeNode) [][]int {
	if root == nil {
		return nil
	}
	level := 0
	q := []*structures.TreeNode{root}
	result := [][]int{}

	for len(q) > 0 {
		result = append(result, []int{})
		levelLength := len(q)
		for i := 0; i < levelLength; i++ {
			node := q[0]
			q = q[1:]
			result[level] = append(result[level], node.Val)
			if node.Left != nil {
				q = append(q, node.Left)
			}
			if node.Right != nil {
				q = append(q, node.Right)
			}
		}
		level++
	}
	return result
}
