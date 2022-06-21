package leetcode

import "testing"

func TestNumSquares(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{12, 3},
		{13, 2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumSquares(tt.n); got != tt.want {
				t.Errorf("NumSquares() = %v, want %v", got, tt.want)
			}
		})
	}
}
