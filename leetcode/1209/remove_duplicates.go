package leetcode

import "strings"

func RemoveDuplicates(s string, k int) string {
	ans := s
	length := -1
	for length != len(ans) {
		length = len(ans)
		for i, count := 0, 1; i < len(ans); i++ {
			if i == 0 || ans[i] != ans[i-1] {
				count = 1
			} else {
				count++
			}
			if count == k {
				ans = ans[:i-k+1] + ans[i+1:]
				break
			}
		}
	}
	return ans
}

func RemoveDuplicatesStack(s string, k int) string {
	var sb strings.Builder
	stack := []rune{}
	count := []int{}
	for _, char := range s {
		if len(stack) == 0 {
			stack = append(stack, char)
			count = append(count, 1)
			continue
		}
		if stack[len(stack)-1] != char {
			stack = append(stack, char)
			count = append(count, 1)
			continue
		}
		if count[len(count)-1] == k-1 {
			count = count[:len(count)-1]
			stack = stack[:len(stack)-1]
		} else {
			count[len(count)-1]++
		}
	}
	for i, _ := range stack {
		for j := 0; j < count[i]; j++ {
			sb.WriteString(string(stack[i]))
		}
	}
	return sb.String()
}
