package leetcode

import (
	"reflect"
	"testing"
)

func TestFourSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   [][]int
	}{
		{[]int{1, 0, -1, 0, -2, 2}, 0, [][]int{{-2, -1, 1, 2}, {-2, 0, 0, 2}, {-1, 0, 0, 1}}},
		{[]int{2, 2, 2, 2, 2}, 8, [][]int{{2, 2, 2, 2}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FourSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FourSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
