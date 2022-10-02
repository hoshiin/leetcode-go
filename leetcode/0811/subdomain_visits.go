package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func SubdomainVisits(cpdomains []string) []string {
	counts := map[string]int{}
	for _, domain := range cpdomains {
		cpinfo := strings.Split(domain, " ")
		frags := strings.Split(cpinfo[1], ".")
		count, _ := strconv.Atoi(cpinfo[0])
		cur := ""
		for i := len(frags) - 1; i >= 0; i-- {
			if i == len(frags)-1 {
				cur = frags[i]
			} else {
				cur = frags[i] + "." + cur
			}
			counts[cur] += count
		}
	}
	ans := make([]string, 0, len(counts))
	for domain, count := range counts {
		str := fmt.Sprintf("%s %s", strconv.Itoa(count), domain)
		ans = append(ans, str)
	}
	return ans
}
