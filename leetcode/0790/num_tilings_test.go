package leetcode

import "testing"

func TestNumTilings(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 5},
		{1, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumTilings(tt.n); got != tt.want {
				t.Errorf("NumTilings() = %v, want %v", got, tt.want)
			}
		})
	}
}
