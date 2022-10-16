package leetcode

import "testing"

func TestGetSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetSum(tt.a, tt.b); got != tt.want {
				t.Errorf("GetSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
