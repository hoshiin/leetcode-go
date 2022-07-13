package leetcode

import "testing"

func TestAddBinary(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{"11", "1", "100"},
		{"1010", "1011", "10101"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := AddBinary(tt.a, tt.b); got != tt.want {
				t.Errorf("AddBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
