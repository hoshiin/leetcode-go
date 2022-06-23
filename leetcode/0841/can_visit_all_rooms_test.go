package leetcode

import "testing"

func TestCanVisitAllRooms(t *testing.T) {
	tests := []struct {
		rooms [][]int
		want  bool
	}{
		{[][]int{{1}, {2}, {3}, {}}, true},
		{[][]int{{1, 3}, {3, 0, 1}, {2}, {0}}, false},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := CanVisitAllRooms(tt.rooms); got != tt.want {
				t.Errorf("CanVisitAllRooms() = %v, want %v", got, tt.want)
			}
		})
	}
}
