package leetcode

import "testing"

func TestFindTargetSumWays(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{1, 1, 1, 1, 1}, 3, 5},
		{[]int{1}, 1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindTargetSumWays(tt.nums, tt.target); got != tt.want {
				t.Errorf("FindTargetSumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
