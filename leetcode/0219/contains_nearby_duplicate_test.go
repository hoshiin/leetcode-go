package leetcode

import "testing"

func TestContainsNearbyDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{1, 2, 3, 1}, 3, true},
		{[]int{1, 0, 1, 1}, 1, true},
		{[]int{1, 2, 3, 1, 2, 3}, 2, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ContainsNearbyDuplicate(tt.nums, tt.k); got != tt.want {
				t.Errorf("ContainsNearbyDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
