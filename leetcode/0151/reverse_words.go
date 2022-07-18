package leetcode

import "fmt"

func ReverseWords(s string) string {
	ans := []string{}
	word := ""
	for _, r := range s {
		if r == ' ' && word != "" {
			ans = append(ans, word)
			word = ""
		}
		if r != ' ' {
			word = word + string(r)
		}
	}
	if word != "" {
		ans = append(ans, word)
	}
	text := ans[len(ans)-1]
	for i := len(ans) - 2; i >= 0; i-- {
		text = fmt.Sprintf("%s %s", text, ans[i])
	}
	return text
}

func ReverseWordsUseIndex(s string) string {
	i, n, res := 0, len(s), ""
	for i < n {
		for i < n && s[i] == ' ' {
			i++
		}
		if i == n {
			break
		}
		j := i
		for j < n && s[j] != ' ' {
			j++
		}
		if len(res) == 0 {
			res = s[i:j]
		} else {
			res = s[i:j] + " " + res
		}
		i = j + 1
	}
	return res
}
