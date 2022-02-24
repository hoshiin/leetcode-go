package leetcode

import "testing"

func TestContainsDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 1}, true},
		{[]int{1, 2, 3, 4}, false},
		{[]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}
	for _, tt := range tests {
		if got := ContainsDuplicate(tt.nums); got != tt.want {
			t.Errorf("ContainsDuplicate() = %v, want %v", got, tt.want)
		}
	}
}
