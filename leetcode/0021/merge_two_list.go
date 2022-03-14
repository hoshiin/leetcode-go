package leetcode

import "github.com/hoshiin/leetcode-go/structures"

func MergeTwoLists(list1 *structures.ListNode, list2 *structures.ListNode) *structures.ListNode {
	if list1 == nil {
		return list2
	} else if list2 == nil {
		return list1
	} else if list1.Val < list2.Val {
		list1.Next = MergeTwoLists(list1.Next, list2)
		return list1
	} else {
		list2.Next = MergeTwoLists(list1, list2.Next)
		return list2
	}
}
