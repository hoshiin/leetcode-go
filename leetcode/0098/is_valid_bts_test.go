package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestIsValidBST(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want bool
	}{
		{structures.NewTreeNode([]int{2, 1, 3}), true},
		{structures.NewTreeNode([]int{5, 1, 4, structures.NULL, structures.NULL, 3, 6}), false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsValidBST(tt.root); got != tt.want {
				t.Errorf("IsValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
