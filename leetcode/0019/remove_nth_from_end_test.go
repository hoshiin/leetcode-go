package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := []struct {
		head *structures.ListNode
		n    int
		want *structures.ListNode
	}{
		{structures.NewListNode([]int{1, 2, 3, 4, 5}), 2, structures.NewListNode([]int{1, 2, 3, 5})},
		{structures.NewListNode([]int{1}), 1, structures.NewListNode([]int{})},
		{structures.NewListNode([]int{1, 2}), 1, structures.NewListNode([]int{1})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
