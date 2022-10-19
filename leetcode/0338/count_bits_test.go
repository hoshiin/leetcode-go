package leetcode

import (
	"reflect"
	"testing"
)

func TestCountBits(t *testing.T) {
	tests := []struct {
		n    int
		want []int
	}{
		{2, []int{0, 1, 1}},
		{5, []int{0, 1, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CountBits(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
