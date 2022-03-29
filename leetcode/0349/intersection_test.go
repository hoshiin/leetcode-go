package leetcode

import (
	"reflect"
	"testing"
)

func TestIntersection(t *testing.T) {
	tests := []struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{9, 4}},
	}
	for _, tt := range tests {
		if got := Intersection(tt.nums1, tt.nums2); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("Intersection() = %v, want %v", got, tt.want)
		}
	}
}
