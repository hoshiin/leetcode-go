package leetcode

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		target int
		nums   []int
		want   int
	}{
		{7, []int{2, 3, 1, 2, 4, 3}, 2},
		{4, []int{1, 4, 4}, 1},
		{11, []int{1, 1, 1, 1, 1, 1, 1, 1}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MinSubArrayLen(tt.target, tt.nums); got != tt.want {
				t.Errorf("MinSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
