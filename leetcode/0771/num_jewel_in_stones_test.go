package leetcode

import "testing"

func TestNumJewelsInStones(t *testing.T) {
	tests := []struct {
		jewels string
		stones string
		want   int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumJewelsInStones(tt.jewels, tt.stones); got != tt.want {
				t.Errorf("NumJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
