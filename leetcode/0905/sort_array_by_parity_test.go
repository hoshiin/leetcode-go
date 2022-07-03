package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArrayByParity(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{3, 1, 2, 4}, []int{2, 4, 3, 1}},
		{[]int{0}, []int{0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SortArrayByParity(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArrayByParity() = %v, want %v", got, tt.want)
			}
		})
	}
}
