package leetcode

import "testing"

func TestMaximalSquare(t *testing.T) {
	tests := []struct {
		matrix [][]byte
		want   int
	}{
		{[][]byte{
			{'1', '0', '1', '0', '0'},
			{'1', '0', '1', '1', '1'},
			{'1', '1', '1', '1', '1'},
			{'1', '0', '0', '1', '0'},
		}, 4},
		{[][]byte{
			{'0', '1'},
			{'1', '0'},
		}, 1},
		{[][]byte{{'0'}}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MaximalSquare(tt.matrix); got != tt.want {
				t.Errorf("MaximalSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
