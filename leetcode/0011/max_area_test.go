package leetcode

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		height []int
		want   int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
		{[]int{1, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxArea(tt.height); got != tt.want {
				t.Errorf("MaxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
