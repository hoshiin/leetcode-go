package leetcode

import "testing"

func TestArrayPairSum(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 4, 3, 2}, 4},
		{[]int{6, 2, 6, 5, 1, 2}, 9},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ArrayPairSum(tt.nums); got != tt.want {
				t.Errorf("ArrayPairSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
