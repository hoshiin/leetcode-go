package leetcode

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		matrix [][]int
		k      int
		want   int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
		{[][]int{{-5}}, 1, -5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := KthSmallest(tt.matrix, tt.k); got != tt.want {
				t.Errorf("KthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
