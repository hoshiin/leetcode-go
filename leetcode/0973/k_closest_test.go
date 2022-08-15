package leetcode

import (
	"reflect"
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		k      int
		want   [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, 1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := KClosest(tt.points, tt.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
