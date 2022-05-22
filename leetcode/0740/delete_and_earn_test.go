package leetcode

import "testing"

func TestDeleteAndEarn(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 4, 2}, 6},
		{[]int{2, 2, 3, 3, 3, 4}, 9},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := DeleteAndEarn(tt.nums); got != tt.want {
				t.Errorf("DeleteAndEarn() = %v, want %v", got, tt.want)
			}
		})
	}
}
