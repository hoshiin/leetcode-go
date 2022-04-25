package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestPostorderTraversal(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want []int
	}{
		// {structures.NewTreeNode([]int{1, structures.NULL, 2, 3}), []int{3, 2, 1}},
		{structures.NewTreeNode([]int{3, 1, 2}), []int{1, 2, 3}},
		// {structures.NewTreeNode([]int{1}), []int{1}},
		// {structures.NewTreeNode([]int{}), nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := PostorderTraversal(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PostorderTraversal() = %v, want %v", got, tt.want)
			}
		})
	}
}
