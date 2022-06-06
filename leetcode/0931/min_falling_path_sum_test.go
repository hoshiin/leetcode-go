package leetcode

import "testing"

func TestMinFallingPathSum(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   int
	}{
		{[][]int{{2, 1, 3}, {6, 5, 4}, {7, 8, 9}}, 13},
		{[][]int{{-19, 57}, {-40, -5}}, -59},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinFallingPathSum(tt.matrix); got != tt.want {
				t.Errorf("MinFallingPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
