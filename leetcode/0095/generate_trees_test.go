package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestGenerateTrees(t *testing.T) {
	tests := []struct {
		n    int
		want []*structures.TreeNode
	}{
		{3, []*structures.TreeNode{structures.NewTreeNode([]int{1, structures.NULL, 2, structures.NULL, 3}),
			structures.NewTreeNode([]int{1, structures.NULL, 3, 2}),
			structures.NewTreeNode([]int{2, 1, 3}),
			structures.NewTreeNode([]int{3, 1, structures.NULL, structures.NULL, 2}),
			structures.NewTreeNode([]int{3, 2, structures.NULL, 1}),
		}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GenerateTrees(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateTrees() = %v, want %v", got, tt.want)
			}
		})
	}
}
