package leetcode

import "fmt"

func RomanToInt(s string) int {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	pairs := map[rune]rune{
		'V': 'I',
		'X': 'I',
		'L': 'X',
		'C': 'X',
		'D': 'C',
		'M': 'C',
	}
	pairNums := map[rune]int{
		'V': 4,
		'X': 9,
		'L': 40,
		'C': 90,
		'D': 400,
		'M': 900,
	}
	beforeChars := map[rune]struct{}{
		'I': {},
		'X': {},
		'C': {},
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		if _, ok := beforeChars[r]; ok && i+1 < len(s) && pairs[rune(s[i+1])] == r {
			sum += pairNums[rune(s[i+1])]
			i++
			continue
		}
		sum += nums[r]
	}
	return sum
}

func RomanToIntOptimized(s string) int {
	values := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}
	sum := 0
	for i := 0; i < len(s); i++ {
		var str string
		if i < len(s)-1 {
			str = s[i : i+2]
		} else {
			str = s[i:]
		}

		if val, ok := values[str]; ok {
			sum += val
			i++
			continue
		}
		sum += values[string(s[i])]
	}
	return sum
}

func RomanToIntRightToLeft(s string) int {
	nums := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sum := nums[rune(s[len(s)-1])]
	for i := len(s) - 2; i >= 0; i-- {
		fmt.Println(string(s[i]))
		if nums[rune(s[i])] < nums[rune(s[i+1])] {
			sum -= nums[rune(s[i])]
		} else {
			sum += nums[rune(s[i])]
		}
	}
	return sum
}
