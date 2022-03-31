package leetcode

import "testing"

func TestTitleToNumber(t *testing.T) {
	tests := []struct {
		columnTitle string
		want        int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}
	for _, tt := range tests {
		if got := TitleToNumber(tt.columnTitle); got != tt.want {
			t.Errorf("TitleToNumber() = %v, want %v", got, tt.want)
		}
	}
}
