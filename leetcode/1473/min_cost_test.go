package leetcode

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		houses []int
		cost   [][]int
		m      int
		n      int
		target int
		want   int
	}{
		{[]int{0, 0, 0, 0, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3, 9},
		{[]int{0, 2, 1, 2, 0}, [][]int{{1, 10}, {10, 1}, {10, 1}, {1, 10}, {5, 1}}, 5, 2, 3, 11},
		{[]int{3, 1, 2, 3}, [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}, {1, 1, 1}}, 4, 3, 3, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinCost(tt.houses, tt.cost, tt.m, tt.n, tt.target); got != tt.want {
				t.Errorf("MinCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
