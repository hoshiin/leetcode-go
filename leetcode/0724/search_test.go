package leetcode

import "testing"

func TestSearch(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{-1, 0, 3, 5, 9, 12}, 9, 4},
		{[]int{-1, 0, 3, 5, 9, 12}, 2, -1},
		{[]int{2, 5}, 2, 0},
	}
	for _, tt := range tests {
		if got := Search(tt.nums, tt.target); got != tt.want {
			t.Errorf("Search() = %v, want %v", got, tt.want)
		}
	}
}
