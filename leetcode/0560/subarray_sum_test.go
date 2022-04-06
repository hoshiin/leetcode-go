package leetcode

import "testing"

func TestSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 1, 1}, 2, 2},
		{[]int{1, 2, 3}, 3, 2},
	}
	for _, tt := range tests {
		if got := SubarraySum(tt.nums, tt.k); got != tt.want {
			t.Errorf("SubarraySum() = %v, want %v", got, tt.want)
		}
	}
}
