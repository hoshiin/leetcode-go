package leetcode

import "testing"

func TestNumDecodings(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"12", 2},
		{"226", 3},
		{"06", 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumDecodings(tt.s); got != tt.want {
				t.Errorf("NumDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
