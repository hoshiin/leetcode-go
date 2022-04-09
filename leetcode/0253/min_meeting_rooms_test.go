package leetcode

import "testing"

func TestMinMeetingRooms(t *testing.T) {
	tests := []struct {
		intervals [][]int
		want      int
	}{
		{[][]int{{0, 30}, {5, 10}, {15, 20}}, 2},
		{[][]int{{7, 10}, {2, 4}}, 1},
	}
	for _, tt := range tests {
		if got := MinMeetingRooms(tt.intervals); got != tt.want {
			t.Errorf("MinMeetingRooms() = %v, want %v", got, tt.want)
		}
	}
}
