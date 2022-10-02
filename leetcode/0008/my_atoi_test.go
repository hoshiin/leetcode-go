package leetcode

import "testing"

func TestMyAtoi(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"42", 42},
		{"-42", -42},
		{"4193 with words", 4193},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MyAtoi(tt.s); got != tt.want {
				t.Errorf("MyAtoi() = %v, want %v", got, tt.want)
			}
		})
	}
}
