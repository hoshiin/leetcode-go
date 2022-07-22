package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestMiddleNode(t *testing.T) {
	tests := []struct {
		head *structures.ListNode
		want *structures.ListNode
	}{
		{structures.NewListNode([]int{1, 2, 3, 4, 5}), structures.NewListNode([]int{3, 4, 5})},
		{structures.NewListNode([]int{1, 2, 3, 4, 5, 6}), structures.NewListNode([]int{4, 5, 6})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MiddleNode(tt.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiddleNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
