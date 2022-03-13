package leetcode

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"1", "2", "Fizz"}},
		{5, []string{"1", "2", "Fizz", "4", "Buzz"}},
		{15, []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"}},
	}
	for _, tt := range tests {
		if got := FizzBuzz(tt.n); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("FizzBuzz() = %v, want %v", got, tt.want)
		}
	}
}
