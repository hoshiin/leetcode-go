package leetcode

import "testing"

func TestThirdMax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3, 2, 1}, 1},
		{[]int{1, 2}, 2},
		{[]int{2, 2, 3, 1}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ThirdMax(tt.nums); got != tt.want {
				t.Errorf("ThirdMax() = %v, want %v", got, tt.want)
			}
		})
	}
}
