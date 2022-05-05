package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestLowestCommonAncestor(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		p    *structures.TreeNode
		q    *structures.TreeNode
		want *structures.TreeNode
	}{
		{structures.NewTreeNode([]int{}), structures.NewTreeNode([]int{}), structures.NewTreeNode([]int{}), structures.NewTreeNode([]int{})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LowestCommonAncestor(tt.root, tt.p, tt.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
