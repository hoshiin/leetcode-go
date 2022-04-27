package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestMaxDepth(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want int
	}{
		{structures.NewTreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}), 3},
		{structures.NewTreeNode([]int{1, structures.NULL, 2}), 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxDepth(tt.root); got != tt.want {
				t.Errorf("MaxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}
