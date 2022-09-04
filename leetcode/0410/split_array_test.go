package leetcode

import "testing"

func TestSplitArray(t *testing.T) {
	tests := []struct {
		nums []int
		m    int
		want int
	}{
		{[]int{7, 2, 5, 10, 8}, 2, 18},
		{[]int{1, 2, 3, 4, 5}, 2, 9},
		{[]int{1, 4, 4}, 3, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SplitArray(tt.nums, tt.m); got != tt.want {
				t.Errorf("SplitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
