package leetcode

import "testing"

func TestReformatDate(t *testing.T) {
	tests := []struct {
		date string
		want string
	}{
		{"20th Oct 2052", "2052-10-20"},
		{"6th Jun 1933", "1933-06-06"},
		{"26th May 1960", "1960-05-26"},
	}
	for _, tt := range tests {
		if got := ReformatDate(tt.date); got != tt.want {
			t.Errorf("ReformatDate() = %v, want %v", got, tt.want)
		}
	}
}
