package leetcode

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 2},
		{3, 3},
	}
	for _, tt := range tests {
		if got := ClimbStairs(tt.n); got != tt.want {
			t.Errorf("ClimbStairs() = %v, want %v", got, tt.want)
		}
	}
}
