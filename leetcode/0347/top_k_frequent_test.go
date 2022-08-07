package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want []int
	}{
		{nums: []int{1, 1, 1, 2, 2, 3}, k: 2, want: []int{1, 2}},
		{nums: []int{1}, k: 1, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TopKFrequent(tt.nums, tt.k); !assert.ElementsMatch(t, got, tt.want) {
				t.Errorf("TopKFrequent() = %v, want %v", got, tt.want)
			}
		})
	}
}
