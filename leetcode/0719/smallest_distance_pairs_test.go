package leetcode

import "testing"

func TestSmallestDistancePair(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{1, 3, 1}, 1, 0},
		{[]int{1, 1, 1}, 2, 0},
		{[]int{1, 6, 1}, 3, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SmallestDistancePair(tt.nums, tt.k); got != tt.want {
				t.Errorf("SmallestDistancePair() = %v, want %v", got, tt.want)
			}
		})
	}
}
