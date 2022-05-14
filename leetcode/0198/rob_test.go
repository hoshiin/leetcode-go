package leetcode

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 1}, 4},
		{[]int{2, 7, 9, 3, 1}, 12},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := Rob(tt.nums); got != tt.want {
				t.Errorf("Rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
