package leetcode

import "testing"

func TestCanConstruct(t *testing.T) {
	tests := []struct {
		ransomNote string
		magazine   string
		want       bool
	}{
		{"a", "b", false},
		{"aa", "ab", false},
		{"aa", "aab", true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CanConstruct(tt.ransomNote, tt.magazine); got != tt.want {
				t.Errorf("CanConstruct() = %v, want %v", got, tt.want)
			}
		})
	}
}
