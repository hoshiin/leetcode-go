package leetcode

import "testing"

func TestConnectSticks(t *testing.T) {
	tests := []struct {
		sticks []int
		want   int
	}{
		{[]int{2, 4, 3}, 14},
		{[]int{1, 8, 3, 5}, 30},
		{[]int{5}, 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ConnectSticks(tt.sticks); got != tt.want {
				t.Errorf("ConnectSticks() = %v, want %v", got, tt.want)
			}
		})
	}
}
