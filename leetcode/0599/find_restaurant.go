package leetcode

import "math"

func FindRestaurant(list1 []string, list2 []string) []string {
	favorites := map[string]int{}
	for i, r := range list1 {
		favorites[r] = i
	}

	minSum, ans := math.MaxInt, []string{}
	for i, r := range list2 {
		if j, ok := favorites[r]; ok {
			if minSum > i+j {
				minSum = i + j
				ans = []string{r}
			} else if minSum == i+j {
				ans = append(ans, r)
			}
		}
	}
	return ans
}
