package leetcode

import "testing"

func TestNextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NextGreatestLetter(tt.letters, tt.target); got != tt.want {
				t.Errorf("NextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
