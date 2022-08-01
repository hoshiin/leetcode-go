package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
	"github.com/stretchr/testify/assert"
)

func TestFindDuplicateSubtrees(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want []*structures.TreeNode
	}{
		{structures.NewTreeNode([]int{1, 2, 3, 4, structures.NULL, 2, 4, structures.NULL, structures.NULL, 4}), []*structures.TreeNode{structures.NewTreeNode([]int{2, 4}), structures.NewTreeNode([]int{4})}},
		{structures.NewTreeNode([]int{2, 1, 1}), []*structures.TreeNode{structures.NewTreeNode([]int{1})}},
		{structures.NewTreeNode([]int{2, 2, 2, 3, structures.NULL, 3, structures.NULL}), []*structures.TreeNode{structures.NewTreeNode([]int{2, 3}), structures.NewTreeNode([]int{3})}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := FindDuplicateSubtrees(tt.root)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
