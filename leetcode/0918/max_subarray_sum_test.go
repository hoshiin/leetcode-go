package leetcode

import "testing"

func TestMaxSubarraySumCircular(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, -2, 3, -2}, 3},
		{[]int{5, -3, 5}, 10},
		{[]int{-3, -2, -3}, -2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxSubarraySumCircular(tt.nums); got != tt.want {
				t.Errorf("MaxSubarraySumCircular() = %v, want %v", got, tt.want)
			}
		})
	}
}
