package leetcode

import "testing"

func TestTribonacci(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{4, 4},
		{25, 1389537},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Tribonacci(tt.n); got != tt.want {
				t.Errorf("Tribonacci() = %v, want %v", got, tt.want)
			}
		})
	}
}
