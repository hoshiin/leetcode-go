package leetcode

import "testing"

func TestToHex(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{26, "1a"},
		{-1, "ffffffff"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ToHex(tt.num); got != tt.want {
				t.Errorf("ToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
