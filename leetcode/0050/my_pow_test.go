package leetcode

import "testing"

func TestMyPow(t *testing.T) {
	tests := []struct {
		x    float64
		n    int
		want float64
	}{
		{2.00000, 10, 1024.00000},
		{2.00000, -2, 0.25000},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MyPow(tt.x, tt.n); got != tt.want {
				t.Errorf("MyPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
