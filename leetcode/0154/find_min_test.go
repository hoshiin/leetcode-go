package leetcode

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 5}, 1},
		{[]int{2, 2, 2, 0, 1}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindMin(tt.nums); got != tt.want {
				t.Errorf("FindMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
