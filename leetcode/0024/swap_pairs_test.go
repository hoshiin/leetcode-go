package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestSwapPairs(t *testing.T) {
	tests := []struct {
		head *structures.ListNode
		want *structures.ListNode
	}{
		{structures.NewListNode([]int{1, 2, 3, 4}), structures.NewListNode([]int{2, 1, 4, 3})},
		{structures.NewListNode([]int{}), structures.NewListNode([]int{})},
		{structures.NewListNode([]int{1}), structures.NewListNode([]int{1})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SwapPairs(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SwapPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
