package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func SwapPairs(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	firstNode := head
	secondNode := head.Next
	firstNode.Next = SwapPairs(secondNode.Next)
	secondNode.Next = firstNode
	return secondNode
}
