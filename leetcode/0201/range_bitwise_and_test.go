package leetcode

import "testing"

func TestRangeBitwiseAnd(t *testing.T) {
	tests := []struct {
		left  int
		right int
		want  int
	}{
		{5, 7, 4},
		{0, 0, 0},
		{0, 2147483647, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RangeBitwiseAnd(tt.left, tt.right); got != tt.want {
				t.Errorf("RangeBitwiseAnd() = %v, want %v", got, tt.want)
			}
		})
	}
}
