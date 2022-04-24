package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestPreorderTraversal(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want []int
	}{
		{structures.NewTreeNode([]int{1, structures.NULL, 2, 3}), []int{1, 2, 3}},
		{structures.NewTreeNode([]int{}), []int{}},
		{structures.NewTreeNode([]int{1}), []int{1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PreorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PreorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
