package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func ReverseList(head *structures.ListNode) *structures.ListNode {
	var prev *structures.ListNode
	curr := head
	for curr != nil {
		nextTemp := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextTemp
	}
	return prev
}

func ReverseListRecursive(head *structures.ListNode) *structures.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := ReverseListRecursive(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}
