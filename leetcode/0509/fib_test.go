package leetcode

import "testing"

func TestFib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{2, 1},
		{3, 2},
		{4, 3},
	}
	for _, tt := range tests {
		if got := Fib(tt.n); got != tt.want {
			t.Errorf("Fib() = %v, want %v", got, tt.want)
		}
	}
}
