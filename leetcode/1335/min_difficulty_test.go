package leetcode

import "testing"

func TestMinDifficulty(t *testing.T) {
	tests := []struct {
		jobDifficulty []int
		d             int
		want          int
	}{
		{[]int{6, 5, 4, 3, 2, 1}, 2, 7},
		{[]int{9, 9, 9}, 4, -1},
		{[]int{1, 1, 1}, 3, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinDifficulty(tt.jobDifficulty, tt.d); got != tt.want {
				t.Errorf("MinDifficulty() = %v, want %v", got, tt.want)
			}
		})
	}
}
