package leetcode

import (
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestClosestValue(t *testing.T) {
	tests := []struct {
		root   *structures.TreeNode
		target float64
		want   int
	}{
		{structures.NewTreeNode([]int{4, 2, 5, 1, 3}), 3.714286, 4},
		{structures.NewTreeNode([]int{1}), 4.428571, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ClosestValue(tt.root, tt.target); got != tt.want {
				t.Errorf("ClosestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
