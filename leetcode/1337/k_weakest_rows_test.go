package leetcode

import (
	"reflect"
	"testing"
)

func TestKWeakestRows(t *testing.T) {
	tests := []struct {
		mat  [][]int
		k    int
		want []int
	}{
		{[][]int{{1, 1, 0, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 0, 0, 0}, {1, 1, 0, 0, 0}, {1, 1, 1, 1, 1}}, 3, []int{2, 0, 3}},
		{[][]int{{1, 0, 0, 0}, {1, 1, 1, 1}, {1, 0, 0, 0}, {1, 0, 0, 0}}, 2, []int{0, 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := KWeakestRows(tt.mat, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KWeakestRows() = %v, want %v", got, tt.want)
			}
		})
	}
}
