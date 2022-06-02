package leetcode

import "testing"

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		m    int
		n    int
		want int
	}{
		{3, 7, 28},
		{3, 2, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := UniquePaths(tt.m, tt.n); got != tt.want {
				t.Errorf("UniquePaths() = %v, want %v", got, tt.want)
			}
		})
	}
}
