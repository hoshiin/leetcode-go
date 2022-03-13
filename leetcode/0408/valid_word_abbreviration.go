package leetcode

import "strconv"

func ValidWordAbbreviation(word string, abbr string) bool {
	i, j := 0, 0
	for i < len(word) && j < len(abbr) {
		if word[i] == abbr[j] {
			i++
			j++
			continue
		}
		if abbr[j] == '0' {
			return false
		}
		if !isDigit(rune(abbr[j])) {
			return false
		}

		k := j
		for k < len(abbr) && isDigit(rune(abbr[k])) {
			k++
		}
		count, _ := strconv.Atoi(abbr[j:k])
		i += count
		j = k
	}
	return i == len(word) && j == len(abbr)
}

func isDigit(s rune) bool {
	return '0' <= s && s <= '9'
}

func ValidWordAbbreviationOptimized(word string, abbr string) bool {
	n, i := 0, 0
	for _, v := range abbr {
		if 'a' <= v && v <= 'z' {
			i += n
			if i >= len(word) || v != rune(word[i]) {
				return false
			}
			n = 0
			i++
			continue
		}
		n = n*10 + int(v-'0')
		if n == 0 {
			return false
		}
	}
	return len(word)-i == n
}
