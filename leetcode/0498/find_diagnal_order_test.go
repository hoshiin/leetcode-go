package leetcode

import (
	"reflect"
	"testing"
)

func TestFindDiagonalOrder(t *testing.T) {
	tests := []struct {
		mat  [][]int
		want []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 4, 7, 5, 3, 6, 8, 9}},
		{[][]int{{1, 2}, {3, 4}}, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindDiagonalOrder(tt.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindDiagonalOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
