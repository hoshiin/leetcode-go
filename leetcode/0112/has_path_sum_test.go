package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestHasPathSum(t *testing.T) {
	tests := []struct {
		root      *structures.TreeNode
		targetSum int
		want      bool
	}{
		{structures.NewTreeNode([]int{5, 4, 8, 11, structures.NULL, 13, 4, 7, 2, structures.NULL, structures.NULL, structures.NULL, 1}), 22, true},
		{structures.NewTreeNode([]int{1, 2, 3}), 5, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := HasPathSum(tt.root, tt.targetSum); got != tt.want {
				t.Errorf("HasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
