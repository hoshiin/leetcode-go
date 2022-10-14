package leetcode

import "testing"

func TestHammingDistance(t *testing.T) {
	tests := []struct {
		x    int
		y    int
		want int
	}{
		{1, 4, 2},
		{3, 1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := HammingDistance(tt.x, tt.y); got != tt.want {
				t.Errorf("HammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
