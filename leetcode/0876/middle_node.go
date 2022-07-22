package leetcode

import "github.com/hoshiin/leetcode-go/structures"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func MiddleNode(head *structures.ListNode) *structures.ListNode {
	n := head
	count := 0
	for n != nil {
		n = n.Next
		count++
	}
	curr := 0
	for curr != count/2 {
		head = head.Next
		curr++
	}
	return head
}

func MiddleNodePointers(head *structures.ListNode) *structures.ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
