package leetcode

import "testing"

func TestOriginalDigits(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"owoztneoer", "012"},
		{"fviefuro", "45"},
	}
	for _, tt := range tests {
		if got := OriginalDigits(tt.s); got != tt.want {
			t.Errorf("OriginalDigits() = %v, want %v", got, tt.want)
		}
	}
}
