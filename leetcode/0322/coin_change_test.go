package leetcode

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		coins  []int
		amount int
		want   int
	}{
		{[]int{1, 2, 5}, 11, 3},
		{[]int{2}, 3, -1},
		{[]int{1}, 0, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CoinChange(tt.coins, tt.amount); got != tt.want {
				t.Errorf("CoinChange() = %v, want %v", got, tt.want)
			}
		})
	}
}
