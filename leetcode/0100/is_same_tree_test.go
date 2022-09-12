package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestIsSameTree(t *testing.T) {
	tests := []struct {
		p    *structures.TreeNode
		q    *structures.TreeNode
		want bool
	}{
		{structures.NewTreeNode([]int{1, 2, 3}), structures.NewTreeNode([]int{1, 2, 3}), true},
		{structures.NewTreeNode([]int{1, 2}), structures.NewTreeNode([]int{1, structures.NULL, 2}), false},
		{structures.NewTreeNode([]int{1, 2, 1}), structures.NewTreeNode([]int{1, 1, 2}), false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsSameTree(tt.p, tt.q); got != tt.want {
				t.Errorf("IsSameTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
