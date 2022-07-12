package leetcode

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{[][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := SpiralOrder(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SpiralOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
