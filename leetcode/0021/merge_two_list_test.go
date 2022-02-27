package leetcode

import (
	"reflect"
	"testing"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 *ListNode
		list2 *ListNode
		want  *ListNode
	}{
		{newListNode([]int{1, 2, 4}), newListNode([]int{1, 3, 4}), newListNode([]int{1, 1, 2, 3, 4, 4})},
		{newListNode([]int{}), newListNode([]int{}), newListNode([]int{})},
		{newListNode([]int{}), newListNode([]int{0}), newListNode([]int{0})},
	}
	for _, tt := range tests {
		if got := MergeTwoLists(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
		}
	}
}

func newListNode(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	l := &ListNode{}
	t := l
	for _, v := range nums {
		t.Next = &ListNode{Val: v}
		t = t.Next
	}
	return l.Next
}
