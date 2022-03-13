package leetcode

import "testing"

func TestValidWordAbbreviation(t *testing.T) {
	tests := []struct {
		word string
		abbr string
		want bool
	}{
		{"internationalization", "i12iz4n", true},
		{"apple", "a2e", false},
	}
	for _, tt := range tests {
		if got := ValidWordAbbreviation(tt.word, tt.abbr); got != tt.want {
			t.Errorf("ValidWordAbbreviation() = %v, want %v", got, tt.want)
		}
	}
}
