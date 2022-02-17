package leetcode

import "testing"

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"III", 3},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}
	for _, tt := range tests {
		if got := RomanToInt(tt.s); got != tt.want {
			t.Errorf("RomanToInt() = %v, want %v", got, tt.want)
		}
	}
}
