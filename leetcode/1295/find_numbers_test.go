package leetcode

import "testing"

func TestFindNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindNumbers(tt.nums); got != tt.want {
				t.Errorf("FindNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
