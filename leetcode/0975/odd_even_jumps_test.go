package leetcode

import "testing"

func TestOddEvenJumps(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{10, 13, 12, 14, 15}, 2},
		{[]int{2, 3, 1, 1, 4}, 3},
		{[]int{5, 1, 3, 4, 2}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := OddEvenJumps(tt.arr); got != tt.want {
				t.Errorf("OddEvenJumps() = %v, want %v", got, tt.want)
			}
		})
	}
}
