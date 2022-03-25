package leetcode

import "testing"

func TestFindKthPositive(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		want int
	}{
		{[]int{2, 3, 4, 7, 11}, 5, 9},
		{[]int{1, 2, 3, 4}, 2, 6},
	}
	for _, tt := range tests {
		if got := FindKthPositive(tt.arr, tt.k); got != tt.want {
			t.Errorf("FindKthPositive() = %v, want %v", got, tt.want)
		}
	}
}
