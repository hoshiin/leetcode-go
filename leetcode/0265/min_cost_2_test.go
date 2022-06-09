package leetcode

import "testing"

func TestMinCostII(t *testing.T) {
	tests := []struct {
		costs [][]int
		want  int
	}{
		{[][]int{{1, 5, 3}, {2, 9, 4}}, 5},
		{[][]int{{1, 3}, {2, 4}}, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinCostII(tt.costs); got != tt.want {
				t.Errorf("MinCostII() = %v, want %v", got, tt.want)
			}
		})
	}
}
