package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func RemoveNthFromEnd(head *structures.ListNode, n int) *structures.ListNode {
	dummy := &structures.ListNode{
		Val:  0,
		Next: head,
	}
	length := 0
	first := head

	for first != nil {
		length++
		first = first.Next
	}

	length -= n
	first = dummy
	for length > 0 {
		length--
		first = first.Next
	}
	first.Next = first.Next.Next

	return dummy.Next
}
