package leetcode

import (
	"reflect"
	"testing"
)

func TestFloodFill(t *testing.T) {
	tests := []struct {
		image [][]int
		sr    int
		sc    int
		color int
		want  [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2, [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
		{[][]int{{0, 0, 0}, {0, 0, 0}}, 0, 0, 0, [][]int{{0, 0, 0}, {0, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FloodFill(tt.image, tt.sr, tt.sc, tt.color); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FloodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
