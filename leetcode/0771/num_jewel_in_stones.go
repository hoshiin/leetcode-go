package leetcode

func NumJewelsInStones(jewels string, stones string) int {
	j := map[rune]struct{}{}
	count := 0
	for _, r := range jewels {
		j[r] = struct{}{}
	}
	for _, stone := range stones {
		if _, ok := j[stone]; ok {
			count++
		}
	}
	return count
}
