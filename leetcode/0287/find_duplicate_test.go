package leetcode

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 4, 2, 2}, 2},
		{[]int{3, 1, 3, 4, 2}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindDuplicate(tt.nums); got != tt.want {
				t.Errorf("FindDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
