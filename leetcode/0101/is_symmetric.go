package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func IsSymmetric(root *structures.TreeNode) bool {
	queue := []*structures.TreeNode{root, root}
	for len(queue) > 0 {
		t1 := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		t2 := queue[len(queue)-1]
		queue = queue[:len(queue)-1]
		if t1 == nil && t2 == nil {
			continue
		}
		if t1 == nil || t2 == nil {
			return false
		}
		if t1.Val != t2.Val {
			return false
		}
		queue = append(queue, t1.Right)
		queue = append(queue, t2.Left)
		queue = append(queue, t1.Left)
		queue = append(queue, t2.Right)
	}
	return true
}

func IsSymmetricRecursive(root *structures.TreeNode) bool {
	return isMirror(root, root)
}

func isMirror(t1, t2 *structures.TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}
	if t1 == nil || t2 == nil {
		return false
	}
	return t1.Val == t2.Val &&
		isMirror(t1.Right, t2.Left) &&
		isMirror(t1.Left, t2.Right)
}
