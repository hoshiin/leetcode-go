package leetcode

import "testing"

func TestNumberOfSteps(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{14, 6},
		{8, 4},
		{123, 12},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumberOfSteps(tt.num); got != tt.want {
				t.Errorf("NumberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
