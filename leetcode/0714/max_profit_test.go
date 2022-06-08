package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		prices []int
		fee    int
		want   int
	}{
		{[]int{1, 3, 2, 8, 4, 9}, 2, 8},
		{[]int{1, 3, 7, 5, 10, 3}, 3, 6},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxProfit(tt.prices, tt.fee); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
