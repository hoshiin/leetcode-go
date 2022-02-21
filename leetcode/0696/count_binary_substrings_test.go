package leetcode

import "testing"

func TestCountBinarySubstrings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"00110011", 6},
		{"10101", 4},
	}
	for _, tt := range tests {
		if got := CountBinarySubstrings(tt.s); got != tt.want {
			t.Errorf("CountBinarySubstrings() = %v, want %v", got, tt.want)
		}
	}
}
