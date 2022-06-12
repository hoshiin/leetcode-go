package leetcode

import "testing"

func TestCountVowelPermutation(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 5},
		{2, 10},
		{5, 68},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CountVowelPermutation(tt.n); got != tt.want {
				t.Errorf("CountVowelPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
