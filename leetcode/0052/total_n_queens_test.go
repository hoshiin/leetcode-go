package leetcode

import "testing"

func TestTotalNQueens(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 2},
		{1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TotalNQueens(tt.n); got != tt.want {
				t.Errorf("TotalNQueens() = %v, want %v", got, tt.want)
			}
		})
	}
}
