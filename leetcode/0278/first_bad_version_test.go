package leetcode

import "testing"

func TestFirstBadVersion(t *testing.T) {
	tests := []struct {
		n    int
		v    int
		want int
	}{
		{5, 4, 4},
		{1, 1, 1},
	}
	for _, tt := range tests {
		version = tt.v
		if got := FirstBadVersion(tt.n); got != tt.want {
			t.Errorf("FirstBadVersion() = %v, want %v", got, tt.want)
		}
	}
}
