package leetcode

import "testing"

func TestHeightChecker(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{5, 1, 2, 3, 4}, 5},
		{[]int{1, 2, 3, 4, 5}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := HeightChecker(tt.heights); got != tt.want {
				t.Errorf("HeightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
