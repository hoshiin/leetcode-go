package leetcode

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		x    int
		want int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Reverse(tt.x); got != tt.want {
				t.Errorf("Reverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
