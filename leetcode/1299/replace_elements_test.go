package leetcode

import (
	"reflect"
	"testing"
)

func TestReplaceElements(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
		{[]int{400}, []int{-1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReplaceElements(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReplaceElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
