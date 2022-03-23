package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestReverseList(t *testing.T) {
	tests := []struct {
		head *structures.ListNode
		want *structures.ListNode
	}{
		{structures.NewListNode([]int{1, 2, 3, 4, 5}), structures.NewListNode([]int{5, 4, 3, 2, 1})},
		{structures.NewListNode([]int{1, 2}), structures.NewListNode([]int{2, 1})},
	}
	for _, tt := range tests {
		if got := ReverseList(tt.head); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("ReverseList() = %v, want %v", got, tt.want)
		}
	}
}
