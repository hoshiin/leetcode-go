package leetcode

import "testing"

func TestStrStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		{"hello", "ll", 2},
		{"aaaaa", "bba", -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := StrStr(tt.haystack, tt.needle); got != tt.want {
				t.Errorf("StrStr() = %v, want %v", got, tt.want)
			}
		})
	}
}
