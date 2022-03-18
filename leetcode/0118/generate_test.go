package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	tests := []struct {
		numRows int
		want    [][]int
	}{
		{5, [][]int{{1}, {1, 1}, {1, 2, 1}, {1, 3, 3, 1}, {1, 4, 6, 4, 1}}},
		{1, [][]int{{1}}},
	}
	for _, tt := range tests {
		if got := Generate(tt.numRows); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Generate() = %v, want %v", got, tt.want)
		}
	}
}
