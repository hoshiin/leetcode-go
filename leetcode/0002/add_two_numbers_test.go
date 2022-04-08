package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		l1   *structures.ListNode
		l2   *structures.ListNode
		want *structures.ListNode
	}{
		{structures.NewListNode([]int{2, 4, 3}), structures.NewListNode([]int{5, 6, 4}), structures.NewListNode([]int{7, 0, 8})},
		{structures.NewListNode([]int{0}), structures.NewListNode([]int{0}), structures.NewListNode([]int{0})},
		{structures.NewListNode([]int{9, 9, 9, 9, 9, 9, 9}), structures.NewListNode([]int{9, 9, 9, 9}), structures.NewListNode([]int{8, 9, 9, 9, 0, 0, 0, 1})},
	}
	for _, tt := range tests {
		if got := AddTwoNumbers(tt.l1, tt.l2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("AddTwoNumbers() = %v, want %v", got, tt.want)
		}
	}
}
