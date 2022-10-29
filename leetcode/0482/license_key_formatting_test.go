package leetcode

import "testing"

func TestLicenseKeyFormatting(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"5F3Z-2e-9-w", 4, "5F3Z-2E9W"},
		{"2-5g-3-J", 2, "2-5G-3J"},
	}
	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			if got := LicenseKeyFormatting(tt.s, tt.k); got != tt.want {
				t.Errorf("LicenseKeyFormatting() = %v, want %v", got, tt.want)
			}
		})
	}
}
