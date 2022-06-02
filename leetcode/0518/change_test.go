package leetcode

import "testing"

func TestChange(t *testing.T) {
	tests := []struct {
		amount int
		coins  []int
		want   int
	}{
		{5, []int{1, 2, 5}, 4},
		{3, []int{2}, 0},
		{10, []int{10}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Change(tt.amount, tt.coins); got != tt.want {
				t.Errorf("Change() = %v, want %v", got, tt.want)
			}
		})
	}
}
