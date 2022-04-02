package leetcode

import (
	"reflect"
	"testing"
)

func TestSummaryRanges(t *testing.T) {
	tests := []struct {
		nums []int
		want []string
	}{
		{[]int{0, 1, 2, 4, 5, 7}, []string{"0->2", "4->5", "7"}},
		{[]int{0, 2, 3, 4, 6, 8, 9}, []string{"0", "2->4", "6", "8->9"}},
	}
	for _, tt := range tests {
		if got := SummaryRanges(tt.nums); !reflect.DeepEqual(got, tt.want) {
			t.Errorf("SummaryRanges() = %v, want %v", got, tt.want)
		}
	}
}
