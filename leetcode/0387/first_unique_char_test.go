package leetcode

import "testing"

func TestFirstUniqChar(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"leetcode", 0},
		{"loveleetcode", 2},
		{"aabb", -1},
	}
	for _, tt := range tests {
		if got := FirstUniqChar(tt.s); got != tt.want {
			t.Errorf("FirstUniqChar() = %v, want %v", got, tt.want)
		}
	}
}
