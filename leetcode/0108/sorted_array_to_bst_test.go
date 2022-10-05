package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestSortedArrayToBST(t *testing.T) {
	tests := []struct {
		nums []int
		want *structures.TreeNode
	}{
		{[]int{-10, -3, 0, 5, 9}, structures.NewTreeNode([]int{0, -10, 5, structures.NULL, -3, structures.NULL, 9})},
		{[]int{1, 3}, structures.NewTreeNode([]int{1, structures.NULL, 3})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SortedArrayToBST(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedArrayToBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
