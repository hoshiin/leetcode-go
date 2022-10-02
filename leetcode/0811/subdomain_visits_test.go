package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubdomainVisits(t *testing.T) {
	tests := []struct {
		cpdomains []string
		want      []string
	}{
		{[]string{"9001 discuss.leetcode.com"}, []string{"9001 leetcode.com", "9001 discuss.leetcode.com", "9001 com"}},
		{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			got := SubdomainVisits(tt.cpdomains)
			assert.ElementsMatch(t, got, tt.want)
		})
	}
}
