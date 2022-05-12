package leetcode

import "testing"

func TestKthGrammar(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{1, 1, 0},
		{2, 1, 0},
		{2, 2, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := KthGrammar(tt.n, tt.k); got != tt.want {
				t.Errorf("KthGrammar() = %v, want %v", got, tt.want)
			}
		})
	}
}
