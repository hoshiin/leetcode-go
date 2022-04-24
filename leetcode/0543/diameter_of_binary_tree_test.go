package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want int
	}{
		{structures.NewTreeNode([]int{1, 2, 3, 4, 5}), 3},
		{structures.NewTreeNode([]int{1, 2}), 1},
		{structures.NewTreeNode([]int{4, -7, -3, structures.NULL, structures.NULL, -9, -3, 9, -7, -4, structures.NULL, 6, structures.NULL, -6, -6, structures.NULL, structures.NULL, 0, 6, 5, structures.NULL, 9, structures.NULL, structures.NULL, -1, -4, structures.NULL, structures.NULL, structures.NULL}), 7},
	}
	for _, tt := range tests {
		if got := DiameterOfBinaryTree(tt.root); got != tt.want {
			t.Errorf("DiameterOfBinaryTree() = %v, want %v", got, tt.want)
		}
	}
}
