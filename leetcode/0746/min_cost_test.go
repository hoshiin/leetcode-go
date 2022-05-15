package leetcode

import "testing"

func TestMinCostClimbingStairs(t *testing.T) {
	tests := []struct {
		cost []int
		want int
	}{
		{[]int{10, 15, 20}, 15},
		{[]int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}, 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinCostClimbingStairs(tt.cost); got != tt.want {
				t.Errorf("MinCostClimbingStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
