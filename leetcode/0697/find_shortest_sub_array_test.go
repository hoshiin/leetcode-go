package leetcode

import "testing"

func TestFindShortestSubArray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 2, 3, 1}, 2},
		{[]int{1, 2, 2, 3, 1, 4, 2}, 6},
	}
	for _, tt := range tests {
		if got := FindShortestSubArray(tt.nums); got != tt.want {
			t.Errorf("FindShortestSubArray() = %v, want %v", got, tt.want)
		}
	}
}
