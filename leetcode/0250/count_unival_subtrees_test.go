package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestCountUnivalSubtrees(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want int
	}{
		{structures.NewTreeNode([]int{5, 1, 5, 5, 5, structures.NULL, 5}), 4},
		{structures.NewTreeNode([]int{}), 0},
		{structures.NewTreeNode([]int{5, 5, 5, 5, 5, structures.NULL, 5}), 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CountUnivalSubtrees(tt.root); got != tt.want {
				t.Errorf("CountUnivalSubtrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
