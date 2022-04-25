package leetcode

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		root *structures.TreeNode
		want [][]int
	}{
		{structures.NewTreeNode([]int{3, 9, 20, structures.NULL, structures.NULL, 15, 7}), [][]int{{3}, {9, 20}, {15, 7}}},
		{structures.NewTreeNode([]int{1}), [][]int{{1}}},
		{structures.NewTreeNode([]int{}), nil},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LevelOrder(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LevelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
