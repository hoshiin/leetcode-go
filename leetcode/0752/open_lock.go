package leetcode

import "strconv"

func OpenLock(deadends []string, target string) int {
	dead := make(map[string]struct{})
	for _, s := range deadends {
		dead[s] = struct{}{}
	}
	q := []queue{{"0000", 0}}

	visited := make(map[string]struct{})
	for len(q) > 0 {
		head := q[0]
		q = q[1:]
		if head.num == target {
			return head.depth
		}
		if _, ok := dead[head.num]; ok {
			continue
		}
		for _, nei := range neighbors(head.num) {
			if _, ok := visited[nei]; !ok {
				visited[nei] = struct{}{}
				q = append(q, queue{nei, head.depth + 1})
			}
		}

	}
	return -1
}

type queue struct {
	num   string
	depth int
}

func neighbors(num string) []string {
	ns := make([]string, 0, 8)
	for i := 0; i < 4; i++ {
		x := int(num[i] - '0')
		nums := []int{-1, 1}
		for _, d := range nums {
			n := (x + d) % 10
			if n < 0 {
				ns = append(ns, num[:i]+strconv.Itoa(10+n)+num[i+1:])
			} else {
				ns = append(ns, num[:i]+strconv.Itoa(n)+num[i+1:])
			}
		}
	}
	return ns
}
