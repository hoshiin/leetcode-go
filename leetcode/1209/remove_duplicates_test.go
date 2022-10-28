package leetcode

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"abcd", 2, "abcd"},
		{"deeedbbcccbdaa", 3, "aa"},
		{"pbbcggttciiippooaais", 2, "ps"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RemoveDuplicates(tt.s, tt.k); got != tt.want {
				t.Errorf("RemoveDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
