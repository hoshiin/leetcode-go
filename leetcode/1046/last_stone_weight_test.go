package leetcode

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		stones []int
		want   int
	}{
		{[]int{2, 7, 4, 1, 8, 1}, 1},
		{[]int{1}, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LastStoneWeight(tt.stones); got != tt.want {
				t.Errorf("LastStoneWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
