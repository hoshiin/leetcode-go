package leetcode

import "testing"

func TestIsHappy(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{19, true},
		{2, false},
	}
	for _, tt := range tests {
		if got := IsHappy(tt.n); got != tt.want {
			t.Errorf("IsHappy() = %v, want %v", got, tt.want)
		}
	}
}
