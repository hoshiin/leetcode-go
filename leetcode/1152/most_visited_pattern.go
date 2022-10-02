package leetcode

import (
	"sort"
	"strings"
)

func MostVisitedPattern(username []string, timestamp []int, website []string) []string {
	order := make([]int, len(username))
	for i := range order {
		order[i] = i
	}
	sort.Slice(order, func(i, j int) bool {
		if username[order[i]] != username[order[j]] {
			return username[order[i]] < username[order[j]]
		}
		return timestamp[order[i]] < timestamp[order[j]]
	})
	threeSequence := make(map[string]int)
	threeSequenceUser := make(map[string]string)
	max := 0
	maxThree := ""
	for i := 0; i < len(order); i++ {
		currentUser := username[order[i]]
		for j := i + 1; j < len(order) && currentUser == username[order[j]]; j++ {
			for k := j + 1; k < len(order) && currentUser == username[order[k]]; k++ {
				seq := website[order[i]] + ":" + website[order[j]] + ":" + website[order[k]]
				if threeSequenceUser[seq] == currentUser {
					continue
				} else {
					threeSequenceUser[seq] = currentUser
				}
				threeSequence[seq]++
				if threeSequence[seq] > max || threeSequence[seq] == max && seq < maxThree {
					max = threeSequence[seq]
					maxThree = seq
				}
			}
		}
	}
	return strings.Split(maxThree, ":")
}
