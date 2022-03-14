package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func IsPalindrome(head *structures.ListNode) bool {
	stack := []int{}
	node := head
	for node.Next != nil {
		stack = append(stack, node.Val)
		node = node.Next
	}
	stack = append(stack, node.Val)
	left, right := 0, len(stack)-1
	for left <= right {
		if stack[left] != stack[right] {
			return false
		}
		left++
		right--
	}
	return true
}
