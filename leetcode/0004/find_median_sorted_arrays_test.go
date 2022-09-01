package leetcode

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  float64
	}{
		{[]int{1, 3}, []int{2}, 2},
		{[]int{1, 2}, []int{3, 4}, 2.5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("FindMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
