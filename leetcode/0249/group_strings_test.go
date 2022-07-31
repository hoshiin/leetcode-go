package leetcode

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGroupStrings(t *testing.T) {
	tests := []struct {
		strings []string
		want    [][]string
	}{
		{[]string{"abc", "bcd", "acef", "xyz", "az", "ba", "a", "z"}, [][]string{{"acef"}, {"a", "z"}, {"abc", "bcd", "xyz"}, {"az", "ba"}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := GroupStrings(tt.strings)
			sort.Slice(got, func(i, j int) bool {
				if len(got[i]) != len(got[j]) {
					return len(got[i]) < len(got[j])
				}
				return got[i][0] < got[j][0]
			})
			sort.Slice(tt.want, func(i, j int) bool {
				if len(tt.want[i]) != len(tt.want[j]) {
					return len(tt.want[i]) < len(tt.want[j])
				}
				return tt.want[i][0] < tt.want[j][0]
			})
			for i, want := range tt.want {
				assert.ElementsMatch(t, got[i], want)
			}
			if len(got) != len(tt.want) {
				t.Errorf("GroupStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
