package leetcode

import "testing"

func TestTotalFruit(t *testing.T) {
	tests := []struct {
		fruits []int
		want   int
	}{
		{[]int{1, 2, 1}, 3},
		{[]int{0, 1, 2, 2}, 3},
		{[]int{1, 2, 3, 2, 2}, 4},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := TotalFruit(tt.fruits); got != tt.want {
				t.Errorf("TotalFruit() = %v, want %v", got, tt.want)
			}
		})
	}
}
