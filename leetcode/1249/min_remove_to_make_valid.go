package leetcode

import "strings"

func MinRemoveToMakeValid(s string) string {
	stack := []int{}
	indexToRemove := make(map[int]struct{})
	for i, r := range s {
		if r == ')' {
			if len(stack) == 0 {
				indexToRemove[i] = struct{}{}
				continue
			}
			stack = stack[:len(stack)-1]
		}
		if r == '(' {
			stack = append(stack, i)
		}
	}
	for len(stack) != 0 {
		indexToRemove[stack[len(stack)-1]] = struct{}{}
		stack = stack[:len(stack)-1]
	}
	var answer strings.Builder
	for i, r := range s {
		if _, ok := indexToRemove[i]; !ok {
			answer.WriteRune(r)
		}
	}
	return answer.String()
}
