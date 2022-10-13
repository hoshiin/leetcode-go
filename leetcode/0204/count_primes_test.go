package leetcode

import "testing"

func TestCountPrimes(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{10, 4},
		{0, 0},
		{1, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CountPrimes(tt.n); got != tt.want {
				t.Errorf("CountPrimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
