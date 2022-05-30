package leetcode

import "testing"

func TestNumWays(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{3, 2, 6},
		{1, 1, 1},
		{3, 2, 6},
		{7, 2, 42},
		{3, 1, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumWays(tt.n, tt.k); got != tt.want {
				t.Errorf("NumWays() = %v, want %v", got, tt.want)
			}
		})
	}
}
