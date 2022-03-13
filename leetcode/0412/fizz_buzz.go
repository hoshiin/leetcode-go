package leetcode

import "fmt"

func FizzBuzz(n int) []string {
	arr := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			arr = append(arr, "FizzBuzz")
		case i%3 == 0:
			arr = append(arr, "Fizz")
		case i%5 == 0:
			arr = append(arr, "Buzz")
		default:
			arr = append(arr, fmt.Sprintf("%d", i))
		}
	}
	return arr
}
