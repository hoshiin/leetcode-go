package leetcode

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{2, 1, 5, 6, 2, 3}, 10},
		{[]int{2, 4}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LargestRectangleArea(tt.heights); got != tt.want {
				t.Errorf("LargestRectangleArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
