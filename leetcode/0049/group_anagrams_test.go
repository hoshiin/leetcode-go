package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := GroupAnagrams(tt.strs)
			sort.Slice(got, func(i, j int) bool { return len(got[i]) < len(got[j]) })
			sort.Slice(tt.want, func(i, j int) bool { return len(tt.want[i]) < len(tt.want[j]) })
			for i, want := range tt.want {
				assert.ElementsMatch(t, got[i], want)
			}
			if len(got) != len(tt.want) {
				t.Errorf("GroupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
