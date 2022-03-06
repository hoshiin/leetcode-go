package leetcode

import "testing"

func TestMinStartValue(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{-3, 2, -3, 4, 2}, 5},
		{[]int{1, 2}, 1},
		{[]int{1, -2, -3}, 5},
	}
	for _, tt := range tests {
		if got := MinStartValue(tt.nums); got != tt.want {
			t.Errorf("MinStartValue() = %v, want %v", got, tt.want)
		}
	}
}
