package leetcode

import "testing"

func TestMySqrt(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{4, 2},
		{8, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MySqrt(tt.x); got != tt.want {
				t.Errorf("MySqrt() = %v, want %v", got, tt.want)
			}
		})
	}
}
