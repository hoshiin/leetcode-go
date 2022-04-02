package leetcode

import "testing"

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{2, 2, 1}, 1},
		{[]int{4, 1, 2, 1, 2}, 4},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		if got := SingleNumber(tt.nums); got != tt.want {
			t.Errorf("SingleNumber() = %v, want %v", got, tt.want)
		}
	}
}
