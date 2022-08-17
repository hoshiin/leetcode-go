package leetcode

import "testing"

func TestFurthestBuilding(t *testing.T) {
	tests := []struct {
		heights []int
		bricks  int
		ladders int
		want    int
	}{
		{[]int{4, 2, 7, 6, 9, 14, 12}, 5, 1, 4},
		{[]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2, 7},
		{[]int{14, 3, 19, 3}, 17, 0, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FurthestBuilding(tt.heights, tt.bricks, tt.ladders); got != tt.want {
				t.Errorf("FurthestBuilding() = %v, want %v", got, tt.want)
			}
		})
	}
}
