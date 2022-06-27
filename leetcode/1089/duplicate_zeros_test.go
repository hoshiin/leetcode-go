package leetcode

import (
	"reflect"
	"testing"
)

func TestDuplicateZeros(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{1, 0, 2, 3, 0, 4, 5, 0}, []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			DuplicateZeros(tt.arr)
			if !reflect.DeepEqual(tt.arr, tt.want) {
				t.Errorf("SortedSquares() = %v, want %v", tt.arr, tt.want)
			}
		})
	}
}
