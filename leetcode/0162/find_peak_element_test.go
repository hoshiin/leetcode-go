package leetcode

import "testing"

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 2},
		{[]int{1, 2, 1, 3, 5, 6, 4}, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindPeakElement(tt.nums); got != tt.want {
				t.Errorf("FindPeakElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
