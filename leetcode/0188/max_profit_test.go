package leetcode

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		k      int
		prices []int
		want   int
	}{
		{2, []int{2, 4, 1}, 2},
		{2, []int{3, 2, 6, 5, 0, 3}, 7},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaxProfit(tt.k, tt.prices); got != tt.want {
				t.Errorf("MaxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}
