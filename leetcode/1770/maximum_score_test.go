package leetcode

import "testing"

func TestMaximumScore(t *testing.T) {
	tests := []struct {
		nums        []int
		multipliers []int
		want        int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, 14},
		{[]int{-5, -3, -3, -2, 7, 1}, []int{-10, -5, 3, 4, 6}, 102},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaximumScore(tt.nums, tt.multipliers); got != tt.want {
				t.Errorf("MaximumScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
