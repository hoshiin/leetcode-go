package leetcode

import "bytes"

func GroupStrings(strings []string) [][]string {
	groups := map[string][]string{}
	var buf bytes.Buffer
	for _, str := range strings {
		for j := 0; j < len(str); j++ {
			buf.WriteByte((26 + str[j] - str[0]) % 26)
		}
		key := buf.String()
		buf.Reset()
		if val, ok := groups[key]; ok {
			groups[key] = append(val, str)
		} else {
			groups[key] = []string{str}
		}
	}
	ans := [][]string{}
	for _, arr := range groups {
		ans = append(ans, arr)
	}
	return ans
}
