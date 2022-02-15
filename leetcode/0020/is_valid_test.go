package leetcode

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"{}[]()", true},
		{"([{}])", true},
		{"{[]()", false},
		{"([{])", false},
		{"[", false},
		{"]", false},
	}
	for _, tt := range tests {
		if got := IsValid(tt.s); got != tt.want {
			t.Errorf("IsValid() = %v, want %v", got, tt.want)
		}
	}
}
