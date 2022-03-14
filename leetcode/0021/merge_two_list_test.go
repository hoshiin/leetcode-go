package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestMergeTwoLists(t *testing.T) {
	tests := []struct {
		list1 *structures.ListNode
		list2 *structures.ListNode
		want  *structures.ListNode
	}{
		{structures.NewListNode([]int{1, 2, 4}), structures.NewListNode([]int{1, 3, 4}), structures.NewListNode([]int{1, 1, 2, 3, 4, 4})},
		{structures.NewListNode([]int{}), structures.NewListNode([]int{}), structures.NewListNode([]int{})},
		{structures.NewListNode([]int{}), structures.NewListNode([]int{0}), structures.NewListNode([]int{0})},
	}
	for _, tt := range tests {
		if got := MergeTwoLists(tt.list1, tt.list2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("MergeTwoLists() = %v, want %v", got, tt.want)
		}
	}
}
