package leetcode

import (
	"reflect"
	"testing"
)

func TestGrayCode(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 3, 2}},
		{1, []int{0, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GrayCode(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GrayCode() = %v, want %v", got, tt.want)
			}
		})
	}
}
