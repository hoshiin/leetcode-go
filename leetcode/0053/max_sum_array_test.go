package leetcode

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
		{[]int{1}, 1},
		{[]int{-1}, -1},
		{[]int{5, 4, -1, 7, 8}, 23},
	}
	for _, tt := range tests {
		if got := MaxSubArray(tt.nums); got != tt.want {
			t.Errorf("MaxSubArray() = %v, want %v", got, tt.want)
		}
	}
}
