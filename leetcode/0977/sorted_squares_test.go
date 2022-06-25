package leetcode

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{-4, -1, 0, 3, 10}, []int{0, 1, 9, 16, 100}},
		{[]int{-7, -3, 2, 3, 11}, []int{4, 9, 9, 49, 121}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SortedSquares(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortedSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
