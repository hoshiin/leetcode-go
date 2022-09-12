package leetcode

import "github.com/hoshiin/leetcode-go/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func IsSameTree(p *structures.TreeNode, q *structures.TreeNode) bool {
	if !check(p, q) {
		return false
	}
	deqP, deqQ := []*structures.TreeNode{p}, []*structures.TreeNode{q}
	for len(deqP) > 0 {
		p = deqP[0]
		deqP = deqP[1:]
		q = deqQ[0]
		deqQ = deqQ[1:]
		if !check(p, q) {
			return false
		}
		if p != nil {
			if !check(p.Left, q.Left) {
				return false
			}
			if p.Left != nil {
				deqP = append(deqP, p.Left)
				deqQ = append(deqQ, q.Left)
			}
			if !check(p.Right, q.Right) {
				return false
			}
			if q.Right != nil {
				deqP = append(deqP, p.Right)
				deqQ = append(deqQ, q.Right)
			}
		}

	}
	return true
}

func check(p, q *structures.TreeNode) bool {
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	return p.Val == q.Val
}
