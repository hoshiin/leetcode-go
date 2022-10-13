package leetcode

import "testing"

func TestHammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{00000000000000000000000000001011, 3},
		{00000000000000000000000010000000, 1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := HammingWeight(tt.num); got != tt.want {
				t.Errorf("HammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
