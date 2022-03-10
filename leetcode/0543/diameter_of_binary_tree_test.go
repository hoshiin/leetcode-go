package leetcode

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *TreeNode
		want int
	}{
		{newTreeNode([]int{1, 2, 3, 4, 5}), 3},
		{newTreeNode([]int{1, 2}), 1},
		{newTreeNode([]int{4, -7, -3, NULL, NULL, -9, -3, 9, -7, -4, NULL, 6, NULL, -6, -6, NULL, NULL, 0, 6, 5, NULL, 9, NULL, NULL, -1, -4, NULL, NULL, NULL}), 7},
	}
	for _, tt := range tests {
		if got := DiameterOfBinaryTree(tt.root); got != tt.want {
			t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
		}
	}
}

var NULL = -1 << 63

func newTreeNode(nums []int) *TreeNode {
	n := len(nums)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: nums[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}
