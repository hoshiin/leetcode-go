package leetcode

import "testing"

func TestMinPathSum(t *testing.T) {
	tests := []struct {
		grid [][]int
		want int
	}{
		{[][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}, 7},
		{[][]int{{1, 2, 3}, {4, 5, 6}}, 12},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinPathSum(tt.grid); got != tt.want {
				t.Errorf("MinPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
