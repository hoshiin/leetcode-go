package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestSearchBST(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		val  int
		want *structures.TreeNode
	}{
		{structures.NewTreeNode([]int{4, 2, 7, 1, 3}), 2, structures.NewTreeNode([]int{2, 1, 3})},
		{structures.NewTreeNode([]int{4, 2, 7, 1, 3}), 5, structures.NewTreeNode([]int{})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SearchBST(tt.root, tt.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SearchBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
