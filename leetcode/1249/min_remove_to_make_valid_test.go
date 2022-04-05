package leetcode

import "testing"

func TestMinRemoveToMakeValid(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"lee(t(c)o)de)", "lee(t(c)o)de"},
		{"a)b(c)d", "ab(c)d"},
		{"))((", ""},
	}
	for _, tt := range tests {
		if got := MinRemoveToMakeValid(tt.s); got != tt.want {
			t.Errorf("MinRemoveToMakeValid() = %v, want %v", got, tt.want)
		}
	}
}
