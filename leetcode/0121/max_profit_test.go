package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		want   int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
	}
	for _, tt := range tests {
		if got := MaxProfit(tt.prices); got != tt.want {
			t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
		}
	}
}
