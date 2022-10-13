package leetcode

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	tests := []struct {
		n    int
		want bool
	}{
		{27, true},
		{0, false},
		{1, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := IsPowerOfThree(tt.n); got != tt.want {
				t.Errorf("IsPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
