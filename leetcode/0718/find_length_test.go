package leetcode

import "testing"

func TestFindLength(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  int
	}{
		{[]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}, 3},
		{[]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindLength(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("FindLength() = %v, want %v", got, tt.want)
			}
		})
	}
}
