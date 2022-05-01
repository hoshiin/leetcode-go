package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestBuildTree(t *testing.T) {
	tests := []struct {
		inorder   []int
		postorder []int
		want      *structures.TreeNode
	}{
		{[]int{9, 3, 15, 20, 7}, []int{9, 15, 7, 20, 3}, structures.NewTreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7})},
		{[]int{-1}, []int{-1}, structures.NewTreeNode([]int{-1})},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := BuildTree(tt.inorder, tt.postorder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BuildTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
