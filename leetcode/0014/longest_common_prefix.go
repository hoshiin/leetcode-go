package leetcode

import (
	"fmt"
	"strings"
)

func LongestCommonPrefix(strs []string) string {
	prefix := ""
	first := strs[0]
	if len(strs) < 2 {
		return first
	}
	second := strs[1]

	for i := 0; i < len(first); i++ {
		if len(second)-1 < i || first[i] != second[i] {
			break
		}
		prefix = fmt.Sprintf("%s%s", string(prefix), string(first[i]))
	}

	if len(strs) < 3 || len(prefix) == 0 {
		return prefix
	}
	for i := 2; i < len(strs); i++ {
		for j := 0; j < len(prefix); j++ {
			if len(strs[i]) < j+1 {
				prefix = prefix[:j]
				break
			}
			if prefix[j] != strs[i][j] {
				prefix = prefix[:j]
				break
			}
		}
	}
	return prefix
}

func LongestCommonPrefixHorizontalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for strings.Index(strs[i], prefix) != 0 {
			prefix = prefix[:len(prefix)-1]
			if len(prefix) == 0 {
				return ""
			}
		}
	}
	return prefix
}

func LongestCommonPrefixVerticalScanning(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	for i := 0; i < len(strs[0]); i++ {
		c := strs[0][i]
		for j := 1; j < len(strs); j++ {
			if (i == len(strs[j])) || strs[j][i] != c {
				return strs[0][:i]
			}
		}
	}
	return strs[0]
}
