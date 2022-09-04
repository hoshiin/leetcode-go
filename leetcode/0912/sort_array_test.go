package leetcode

import (
	"reflect"
	"testing"
)

func TestSortArray(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{5, 2, 3, 1}, []int{1, 2, 3, 5}},
		{[]int{5, 1, 1, 2, 0, 0}, []int{0, 0, 1, 1, 2, 5}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SortArray(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SortArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
