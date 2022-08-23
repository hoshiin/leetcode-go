package leetcode

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 5, 1, 2}, 1},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0},
		{[]int{11, 13, 15, 17}, 11},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindMin(tt.nums); got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
