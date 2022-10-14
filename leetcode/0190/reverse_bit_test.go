package leetcode

import "testing"

func TestReverseBits(t *testing.T) {
	tests := []struct {
		num  uint32
		want uint32
	}{
		{43261596, 964176192},
		{4294967293, 3221225471},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ReverseBits(tt.num); got != tt.want {
				t.Errorf("ReverseBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
