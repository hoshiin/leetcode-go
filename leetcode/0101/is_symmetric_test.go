package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestIsSymmetric(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want bool
	}{
		{structures.NewTreeNode([]int{1, 2, 2, 3, 4, 4, 3}), true},
		{structures.NewTreeNode([]int{1, 2, 2, structures.NULL, 3, structures.NULL, 3}), false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsSymmetric(tt.root); got != tt.want {
				t.Errorf("IsSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}
