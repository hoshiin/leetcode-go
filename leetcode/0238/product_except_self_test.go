package leetcode

import (
	"reflect"
	"testing"
)

func TestProductExceptSelf(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{24, 12, 8, 6}},
		{[]int{-1, 1, 0, -3, 3}, []int{0, 0, 9, 0, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ProductExceptSelf(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ProductExceptSelf() = %v, want %v", got, tt.want)
			}
		})
	}
}
