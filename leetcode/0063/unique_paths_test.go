package leetcode

import "testing"

func TestUniquePathsWithObstacles(t *testing.T) {
	tests := []struct {
		obstacleGrid [][]int
		want         int
	}{
		{[][]int{{0, 0, 0}, {0, 1, 0}, {0, 0, 0}}, 2},
		{[][]int{{0, 1}, {0, 0}}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := UniquePathsWithObstacles(tt.obstacleGrid); got != tt.want {
				t.Errorf("UniquePathsWithObstacles() = %v, want %v", got, tt.want)
			}
		})
	}
}
