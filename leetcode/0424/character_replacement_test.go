package leetcode

import "testing"

func TestCharacterReplacement(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"", args{"ABAB", 2}, 4},
		{"", args{"AABABBA", 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CharacterReplacement(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("CharacterReplacement() = %v, want %v", got, tt.want)
			}
		})
	}
}
