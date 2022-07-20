package leetcode

import (
	"reflect"
	"testing"
)

func TestRunningSum(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := RunningSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RunningSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
