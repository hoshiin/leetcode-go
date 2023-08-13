package main

import (
	"reflect"
	"testing"

	"github.com/hoshiin/leetcode-go/structures"
)

func TestInvertTree(t *testing.T) {
	type args struct {
		root *structures.TreeNode
	}
	tests := []struct {
		name string
		args args
		want *structures.TreeNode
	}{
		{"", args{root: structures.NewTreeNode([]int{4, 2, 7, 1, 3, 6, 9})}, structures.NewTreeNode([]int{4, 7, 2, 9, 6, 3, 1})},
		{"", args{root: structures.NewTreeNode([]int{2, 1, 3})}, structures.NewTreeNode([]int{2, 3, 1})},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InvertTree(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InvertTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
