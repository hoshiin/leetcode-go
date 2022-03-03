package leetcode

import "testing"

func TestAddStrings(t *testing.T) {
	tests := []struct {
		num1 string
		num2 string
		want string
	}{
		{"11", "123", "134"},
		{"456", "77", "533"},
		{"0", "0", "0"},
	}
	for _, tt := range tests {
		if got := AddStrings(tt.num1, tt.num2); got != tt.want {
			t.Errorf("AddStrings() = %v, want %v", got, tt.want)
		}
	}
}
