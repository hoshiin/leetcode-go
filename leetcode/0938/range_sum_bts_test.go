package leetcode

import "testing"

func TestRangeSumBST(t *testing.T) {
	tests := []struct {
		root *TreeNode
		low  int
		high int
		want int
	}{
		{newTreeNode([]int{10, 5, 15, 3, 7, NULL, 18}), 7, 15, 32},
		{newTreeNode([]int{10, 5, 15, 3, 7, 13, 18, 1, NULL, 6}), 6, 10, 23},
	}
	for _, tt := range tests {
		if got := RangeSumBST(tt.root, tt.low, tt.high); got != tt.want {
			t.Errorf("RangeSumBST() = %v, want %v", got, tt.want)
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
