package leetcode

import (
	"reflect"
	"testing"
)

func TestGetRow(t *testing.T) {
	tests := []struct {
		rowIndex int
		want     []int
	}{
		{3, []int{1, 3, 3, 1}},
		{0, []int{1}},
		{1, []int{1, 1}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GetRow(tt.rowIndex); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRow() = %v, want %v", got, tt.want)
			}
		})
	}
}
