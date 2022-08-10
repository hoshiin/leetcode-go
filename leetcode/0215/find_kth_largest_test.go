package leetcode

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindKthLargest(tt.nums, tt.k); got != tt.want {
				t.Errorf("FindKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
