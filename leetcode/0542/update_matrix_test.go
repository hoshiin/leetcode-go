package leetcode

import (
	"reflect"
	"testing"
)

func TestUpdateMatrix(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want [][]int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, [][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}},
		{[][]int{{0, 0, 0}, {0, 1, 0}, {1, 1, 1}}, [][]int{{0, 0, 0}, {0, 1, 0}, {1, 2, 1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := UpdateMatrix(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
