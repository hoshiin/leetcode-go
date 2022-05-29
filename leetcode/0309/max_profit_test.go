package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{1, 2, 3, 0, 2}, 3},
		{[]int{1}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxProfit(tt.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
