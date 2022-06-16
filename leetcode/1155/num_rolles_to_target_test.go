package leetcode

import "testing"

func TestNumRollsToTarget(t *testing.T) {
	tests := []struct {
		n      int
		k      int
		target int
		want   int
	}{
		{1, 6, 3, 1},
		{2, 6, 7, 6},
		{30, 30, 500, 222616187},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumRollsToTarget(tt.n, tt.k, tt.target); got != tt.want {
				t.Errorf("NumRollsToTarget() = %v, want %v", got, tt.want)
			}
		})
	}
}
