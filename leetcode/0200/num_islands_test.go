package leetcode

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{[][]byte{
			{'1', '1', '1', '1', '0'},
			{'1', '1', '0', '1', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '0', '0', '0'},
		}, 1},
		{[][]byte{
			{'1', '1', '0', '0', '0'},
			{'1', '1', '0', '0', '0'},
			{'0', '0', '1', '0', '0'},
			{'0', '0', '0', '1', '1'},
		}, 3},
	}
	for _, tt := range tests {
		if got := NumIslands(tt.grid); got != tt.want {
			t.Errorf("NumIslands() = %v, want %v", got, tt.want)
		}
	}
}
