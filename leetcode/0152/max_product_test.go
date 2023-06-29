package leetcode

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 3, -2, 4}, 6},
		{[]int{-2, 0, -1}, 0},
		{[]int{-2, 3, -4}, 24},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxProduct(tt.nums); got != tt.want {
				t.Errorf("MaxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
