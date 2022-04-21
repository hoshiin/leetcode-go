package leetcode

import "testing"

func TestCanPartition(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 5, 11, 5}, true},
		{[]int{1, 2, 3, 5}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CanPartition(tt.nums); got != tt.want {
				t.Errorf("CanPartition() = %v, want %v", got, tt.want)
			}
		})
	}
}
