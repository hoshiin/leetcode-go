package leetcode

import (
	"reflect"
	"testing"
)

func TestMostVisitedPattern(t *testing.T) {
	tests := []struct {
		username  []string
		timestamp []int
		website   []string
		want      []string
	}{
		{[]string{"joe", "joe", "joe", "james", "james", "james", "james", "mary", "mary", "mary"}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []string{"home", "about", "career", "home", "cart", "maps", "home", "home", "about", "career"}, []string{"home", "about", "career"}},
		{[]string{"ua", "ua", "ua", "ub", "ub", "ub"}, []int{1, 2, 3, 4, 5, 6}, []string{"a", "b", "a", "a", "b", "c"}, []string{"a", "b", "a"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := MostVisitedPattern(tt.username, tt.timestamp, tt.website); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MostVisitedPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
