package leetcode

import "testing"

func TestMinCost(t *testing.T) {
	tests := []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{17, 2, 17}, {16, 16, 5}, {14, 3, 19}}, 10},
		{[][]int{{7, 6, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinCost(tt.costs); got != tt.want {
				t.Errorf("MinCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
