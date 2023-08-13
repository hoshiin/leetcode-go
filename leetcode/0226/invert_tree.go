package main

import "github.com/hoshiin/leetcode-go/structures"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func InvertTree(root *structures.TreeNode) *structures.TreeNode {
	if root == nil {
		return nil
	}
	stack := []*structures.TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		node.Right, node.Left = node.Left, node.Right
		if node.Right != nil {
			stack = append(stack, node.Right)
		}
		if node.Left != nil {
			stack = append(stack, node.Left)
		}
	}
	return root
}

func InvertTreeRecursive(root *structures.TreeNode) *structures.TreeNode {
	if root == nil {
		return root
	}
	tmp := InvertTreeRecursive(root.Left)
	root.Left = InvertTreeRecursive(root.Right)
	root.Right = tmp
	return root
}
