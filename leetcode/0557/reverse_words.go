package leetcode

import (
	"fmt"
	"strings"
)

func ReverseWords(s string) string {
	start, end := 0, 0
	ans := []string{}
	for start < len(s) && end < len(s) {
		for s[end] != ' ' {
			end++
			if end >= len(s)-1 {
				break
			}
		}
		if end == len(s)-1 && s[end] != ' ' {
			ans = append(ans, reverse(s[start:]))
			break
		}
		ans = append(ans, reverse(s[start:end]))
		end++
		start = end
	}
	text := ans[0]
	for i := 1; i < len(ans); i++ {
		text = fmt.Sprintf("%s %s", text, ans[i])
	}
	return text
}

func reverse(s string) string {
	start, end := 0, len(s)-1
	runes := []rune(s)
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
	return string(runes)
}

func ReverseWordsStrings(s string) string {
	ss := strings.Split(s, " ")
	for i, s := range ss {
		ss[i] = reverse(s)
	}
	return strings.Join(ss, " ")
}

func ReverseWordsRune(s string) string {
	ans := []rune(s)
	left := 0
	for i, r := range ans {
		if r == ' ' {
			reverseRune(ans, left, i-1)
			left = i + 1
		}
	}
	reverseRune(ans, left, len(ans)-1)
	return string(ans)
}

func reverseRune(runes []rune, start, end int) {
	for start < end {
		runes[start], runes[end] = runes[end], runes[start]
		start++
		end--
	}
}
