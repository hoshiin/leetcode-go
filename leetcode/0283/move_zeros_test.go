package leetcode

import (
	"reflect"
	"testing"
)

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0}, []int{0}},
	}
	for _, tt := range tests {
		MoveZeroes(tt.nums)
		if !reflect.DeepEqual(tt.nums, tt.want) {
			t.Errorf("MoveZeroes() = %v, want %v", tt.nums, tt.want)
		}
	}
}
