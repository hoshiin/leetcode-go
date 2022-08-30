package leetcode

import "testing"

func TestIsPerfectSquare(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{16, true},
		{14, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsPerfectSquare(tt.num); got != tt.want {
				t.Errorf("IsPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
