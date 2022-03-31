package leetcode

import "math"

func TitleToNumber(columnTitle string) int {
	sum := 0.0
	for i := 0; i < len(columnTitle); i++ {
		sum += math.Pow(26.0, float64(len(columnTitle)-i-1)) * float64(toNumber(rune(columnTitle[i])))
	}
	return int(sum)
}

func toNumber(column rune) int {
	return int(column-'A') + 1
}

func TitleToNumberOptimized(columnTitle string) int {
	sum := 0
	for _, c := range columnTitle {
		sum = sum*26 + int(c-'A'+1)
	}
	return sum
}
