package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWallsAndGates(t *testing.T) {
	tests := []struct {
		rooms [][]int
		want  [][]int
	}{
		{[][]int{{2147483647, -1, 0, 2147483647}, {2147483647, 2147483647, 2147483647, -1}, {2147483647, -1, 2147483647, -1}, {0, -1, 2147483647, 2147483647}}, [][]int{{3, -1, 0, 1}, {2, 2, 1, -1}, {1, -1, 2, -1}, {0, -1, 3, 4}}},
		{[][]int{{-1}}, [][]int{{-1}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			WallsAndGates(tt.rooms)
			assert.Equal(t, tt.rooms, tt.want)
		})
	}
}
