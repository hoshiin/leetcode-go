package leetcode

func IsValid(s string) bool {
	stack := []rune{}
	mapping := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	for _, r := range s {
		if b, ok := mapping[r]; ok {
			if len(stack) == 0 || b != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, r)
		}
	}
	return len(stack) == 0
}
