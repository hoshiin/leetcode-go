package leetcode

import "testing"

func TestValidMountainArray(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{2, 1}, false},
		{[]int{3, 5, 5}, false},
		{[]int{0, 3, 2, 1}, true},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ValidMountainArray(tt.arr); got != tt.want {
				t.Errorf("ValidMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
