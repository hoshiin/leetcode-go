package leetcode

import "testing"

func TestDivide(t *testing.T) {
	tests := []struct {
		dividend int
		divisor  int
		want     int
	}{
		{10, 3, 3},
		{7, -3, -2},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Divide(tt.dividend, tt.divisor); got != tt.want {
				t.Errorf("Divide() = %v, want %v", got, tt.want)
			}
		})
	}
}
