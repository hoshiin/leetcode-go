package leetcode

import "testing"

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 0, 1, 1, 0}, 4},
		{[]int{1, 0, 1, 1, 0, 1}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindMaxConsecutiveOnes(tt.nums); got != tt.want {
				t.Errorf("FindMaxConsecutiveOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
