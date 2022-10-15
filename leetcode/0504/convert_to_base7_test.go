package leetcode

import "testing"

func TestConvertToBase7(t *testing.T) {
	tests := []struct {
		num  int
		want string
	}{
		{100, "202"},
		{-7, "-10"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := ConvertToBase7(tt.num); got != tt.want {
				t.Errorf("ConvertToBase7() = %v, want %v", got, tt.want)
			}
		})
	}
}
