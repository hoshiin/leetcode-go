package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleNumber(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 1, 3, 2, 5}, []int{3, 5}},
		{[]int{-1, 0}, []int{-1, 0}},
		{[]int{0, 1}, []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SingleNumber(tt.nums)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
