package leetcode

import "strconv"

func EvalRPN(tokens []string) int {
	stack := []string{}
	for _, token := range tokens {
		if _, err := strconv.Atoi(token); err == nil {
			stack = append(stack, token)
		} else {
			n2, _ := strconv.Atoi(stack[len(stack)-1])
			stack = append(stack[:len(stack)-1])
			n1, _ := strconv.Atoi(stack[len(stack)-1])
			stack = append(stack[:len(stack)-1])
			result := applyOperator(token, n1, n2)
			stack = append(stack, result)
		}
	}
	i, _ := strconv.Atoi(stack[len(stack)-1])
	return i
}

func applyOperator(token string, n1, n2 int) string {
	switch token {
	case "+":
		return strconv.Itoa(n1 + n2)
	case "-":
		return strconv.Itoa(n1 - n2)
	case "*":
		return strconv.Itoa(n1 * n2)
	}
	return strconv.Itoa(n1 / n2)
}
