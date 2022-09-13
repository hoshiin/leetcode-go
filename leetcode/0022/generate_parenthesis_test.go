package leetcode

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	tests := []struct {
		n    int
		want []string
	}{
		{3, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}},
		{1, []string{"()"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := GenerateParenthesis(tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GenerateParenthesis() = %v, want %v", got, tt.want)
			}
		})
	}
}
