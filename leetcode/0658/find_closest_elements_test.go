package leetcode

import (
	"reflect"
	"testing"
)

func TestFindClosestElements(t *testing.T) {
	tests := []struct {
		arr  []int
		k    int
		x    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5}, 4, 3, []int{1, 2, 3, 4}},
		{[]int{1, 2, 3, 4, 5}, 4, -1, []int{1, 2, 3, 4}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := FindClosestElements(tt.arr, tt.k, tt.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindClosestElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
