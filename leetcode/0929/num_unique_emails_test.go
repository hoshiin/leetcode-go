package leetcode

import "testing"

func TestNumUniqueEmails(t *testing.T) {
	tests := []struct {
		emails []string
		want   int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
		{[]string{"a@leetcode.com", "b@leetcode.com", "c@leetcode.com"}, 3},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := NumUniqueEmails(tt.emails); got != tt.want {
				t.Errorf("NumUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
