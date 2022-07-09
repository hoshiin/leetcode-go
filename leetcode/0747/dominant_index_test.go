package leetcode

import "testing"

func TestDominantIndex(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 6, 1, 0}, 1},
		{[]int{1, 2, 3, 4}, -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DominantIndex(tt.nums); got != tt.want {
				t.Errorf("DominantIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
