package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func DeleteNode(node *structures.ListNode) {
	nextNode := node.Next
	node.Val = nextNode.Val
	node.Next = nextNode.Next
}
