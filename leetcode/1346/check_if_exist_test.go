package leetcode

import "testing"

func TestCheckIfExist(t *testing.T) {
	tests := []struct {
		arr  []int
		want bool
	}{
		{[]int{10, 2, 5, 3}, true},
		{[]int{7, 1, 14, 11}, true},
		{[]int{3, 1, 7, 11}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CheckIfExist(tt.arr); got != tt.want {
				t.Errorf("CheckIfExist() = %v, want %v", got, tt.want)
			}
		})
	}
}
