package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func AddTwoNumbers(l1 *structures.ListNode, l2 *structures.ListNode) *structures.ListNode {
	h1, h2 := l1, l2
	carry := 0
	dummy := &structures.ListNode{}
	curr := dummy
	for h1 != nil || h2 != nil {
		sum := 0
		if h1 != nil {
			sum += h1.Val
			h1 = h1.Next
		}
		if h2 != nil {
			sum += h2.Val
			h2 = h2.Next
		}
		curr.Next = &structures.ListNode{Val: (sum + carry) % 10}
		carry = (sum + carry) / 10
		curr = curr.Next
	}
	if carry != 0 {
		curr.Next = &structures.ListNode{Val: carry}
	}
	return dummy.Next
}
