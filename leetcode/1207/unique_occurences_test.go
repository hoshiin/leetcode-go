package leetcode

import "testing"

func TestUniqueOccurrences(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{1, 2, 2, 1, 1, 3}, true},
		{[]int{1, 2}, false},
		{[]int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}, true},
	}
	for _, tt := range tests {
		if got := UniqueOccurrences(tt.arr); got != tt.want {
			t.Errorf("UniqueOccurrences() = %v, want %v", got, tt.want)
		}
	}
}
