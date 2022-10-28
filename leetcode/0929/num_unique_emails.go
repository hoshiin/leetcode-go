package leetcode

import "strings"

func NumUniqueEmails(emails []string) int {
	origins := map[string]struct{}{}
	for _, email := range emails {
		origin := originEmail(email)
		if _, ok := origins[origin]; !ok {
			origins[origin] = struct{}{}
		}
	}
	return len(origins)
}

func originEmail(email string) string {
	var origin strings.Builder
	skip := false
	for i, char := range email {
		if char == '+' {
			skip = true
		}
		if char == '@' {
			origin.WriteString(email[i:])
			break
		}
		if !skip && char != '.' {
			origin.WriteRune(char)
		}
	}
	return origin.String()
}
