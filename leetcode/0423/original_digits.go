package leetcode

import (
	"fmt"
	"strings"
)

func OriginalDigits(s string) string {
	counts := make(map[rune]int)
	for _, letter := range s {
		counts[letter]++
	}
	out := make([]int, 10)
	// "z" is only represented in "zero"
	out[0] = counts['z']
	// "x" is only represented in "six"
	out[6] = counts['x']
	// "h" is only represented in "eight"
	out[8] = counts['g']
	// "w" is only represented in "two"
	out[2] = counts['w']
	// "u" is only represented in "four"
	out[4] = counts['u']
	// "h" is represented only in "three" and "eight"
	out[3] = counts['h'] - out[8]
	// "f" is represented only in "four" and "five"
	out[5] = counts['f'] - out[4]
	// "s" is represented only in "seven" and "six"
	out[7] = counts['s'] - out[6]
	// "o" is represented in "one", "zero", "two", "four"
	out[1] = counts['o'] - out[0] - out[2] - out[4]
	// "i" is represented in "nine", "six", "five", "eight"
	out[9] = counts['i'] - out[6] - out[5] - out[8]

	var b strings.Builder
	for i, n := range out {
		for j := 0; j < n; j++ {
			fmt.Fprintf(&b, "%d", i)
		}
	}
	return b.String()
}
