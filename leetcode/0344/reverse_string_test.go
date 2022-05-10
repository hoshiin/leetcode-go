package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseString(t *testing.T) {
	tests := []struct {
		s    []byte
		want []byte
	}{
		{[]byte{'h', 'e', 'l', 'l', 'o'}, []byte{'o', 'l', 'l', 'e', 'h'}},
		{[]byte{'H', 'a', 'n', 'n', 'a', 'h'}, []byte{'h', 'a', 'n', 'n', 'a', 'H'}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			ReverseString(tt.s)
			assert.EqualValues(t, tt.s, tt.want)
		})
	}
}
