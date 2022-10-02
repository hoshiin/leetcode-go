package leetcode

import "testing"

func TestReorderedPowerOf2(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{1, true},
		{10, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReorderedPowerOf2(tt.n); got != tt.want {
				t.Errorf("ReorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
