package leetcode

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		strs []string
		want string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{"cb", "cbb", "a"}, ""},
	}
	for _, tt := range tests {
		if got := LongestCommonPrefix(tt.strs); got != tt.want {
			t.Errorf("LongestCommonPrefix() = %v, want %v", got, tt.want)
		}
	}
}
