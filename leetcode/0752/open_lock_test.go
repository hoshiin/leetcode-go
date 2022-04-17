package leetcode

import "testing"

func TestOpenLock(t *testing.T) {
	tests := []struct {
		deadends []string
		target   string
		want     int
	}{
		{[]string{"0201", "0101", "0102", "1212", "2002"}, "0202", 6},
		{[]string{"8888"}, "0009", 1},
		{[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888", -1},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := OpenLock(tt.deadends, tt.target); got != tt.want {
				t.Errorf("OpenLock() = %v, want %v", got, tt.want)
			}
		})
	}
}
