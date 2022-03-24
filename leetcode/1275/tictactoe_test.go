package leetcode

import "testing"

func TestTictactoe(t *testing.T) {
	tests := []struct {
		moves [][]int
		want  string
	}{
		{[][]int{{0, 0}, {2, 0}, {1, 1}, {2, 1}, {2, 2}}, "A"},
		{[][]int{{0, 0}, {1, 1}, {0, 1}, {0, 2}, {1, 0}, {2, 0}}, "B"},
		{[][]int{{0, 0}, {1, 1}, {2, 0}, {1, 0}, {1, 2}, {2, 1}, {0, 1}, {0, 2}, {2, 2}}, "Draw"},
	}
	for _, tt := range tests {
		if got := Tictactoe(tt.moves); got != tt.want {
			t.Errorf("Tictactoe() = %v, want %v", got, tt.want)
		}
	}
}
