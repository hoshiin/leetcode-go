package leetcode

import (
	"fmt"
	"strings"
)

func RemoveDuplicates(s string) string {
	duplicates := make(map[string]struct{})
	for i := 'a'; i <= 'z'; i++ {
		duplicates[fmt.Sprintf("%s%s", string(i), string(i))] = struct{}{}
	}
	prevLength := -1
	for prevLength != len(s) {
		prevLength = len(s)
		for d := range duplicates {
			s = strings.Replace(s, d, "", -1)
		}
	}
	return s
}

func RemoveDuplicatesStack(s string) string {
	stack := []string{}
	for _, r := range s {
		if len(stack) == 0 || string(r) != stack[len(stack)-1] {
			stack = append(stack, string(r))
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return strings.Join(stack, "")
}

func RemoveDuplicatesStackOptimized(s string) string {
	stack := make([]byte, 0, len(s))
	for _, b := range []byte(s) {
		if len(stack) > 0 && b == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, b)
		}
	}
	return string(stack)
}

func RemoveDuplicatesStackRangeOptimized(s string) string {
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		if len(stack) > 0 && s[i] == stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return string(stack)
}
