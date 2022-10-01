package leetcode

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3}, []int{1, 3, 2}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 1, 5}, []int{1, 5, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			NextPermutation(tt.nums)
			if !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("NextPermutation() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}
