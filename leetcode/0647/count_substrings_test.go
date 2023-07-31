package leetcode

import "testing"

func TestCountSubstrings(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"abc"}, 3},
		{"", args{"aaa"}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountSubstrings(tt.args.s); got != tt.want {
				t.Errorf("CountSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
