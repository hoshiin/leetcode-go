package leetcode

import "testing"

func TestMaximumUnits(t *testing.T) {
	tests := []struct {
		boxTypes  [][]int
		truckSize int
		want      int
	}{
		{[][]int{{1, 3}, {2, 2}, {3, 1}}, 4, 8},
		{[][]int{{5, 10}, {2, 5}, {4, 7}, {3, 9}}, 10, 91},
	}
	for _, tt := range tests {
		if got := MaximumUnits(tt.boxTypes, tt.truckSize); got != tt.want {
			t.Errorf("MaximumUnits() = %v, want %v", got, tt.want)
		}
	}
}
