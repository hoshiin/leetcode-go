package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"abbaca", "ca"},
		{"azxxzy", "ay"},
	}
	for _, tt := range tests {
		if got := RemoveDuplicates(tt.s); got != tt.want {
			t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
		}
	}
}
