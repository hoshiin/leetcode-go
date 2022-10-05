package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func SortedArrayToBST(nums []int) *structures.TreeNode {
	return helper(nums, 0, len(nums)-1)
}

func helper(nums []int, left, right int) *structures.TreeNode {
	if left > right {
		return nil
	}

	// always choose left middle node as a root
	p := (left + right) / 2

	root := &structures.TreeNode{Val: nums[p]}
	root.Left = helper(nums, left, p-1)
	root.Right = helper(nums, p+1, right)
	return root
}
