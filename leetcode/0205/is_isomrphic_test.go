package leetcode

import "testing"

func TestIsIsomorphic(t *testing.T) {
	tests := []struct {
		s    string
		t    string
		want bool
	}{
		{"egg", "add", true},
		{"foo", "bar", false},
		{"paper", "title", true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsIsomorphic(tt.s, tt.t); got != tt.want {
				t.Errorf("IsIsomorphic() = %v, want %v", got, tt.want)
			}
		})
	}
}
