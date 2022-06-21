package leetcode

func NumSquares(n int) int {
	squares := []int{}
	for i := 1; i*i <= n; i++ {
		squares = append(squares, i*i)
	}

	queue := []int{n}
	seen := make([]bool, n)
	ans := 0

	for len(queue) > 0 {
		ans++

		newQ := []int{}
		for _, pop := range queue {
			for _, square := range squares {
				if square == pop {
					return ans
				}
				reminder := pop - square
				if reminder > 0 && !seen[reminder] {
					newQ = append(newQ, reminder)
					seen[reminder] = true
				}
			}
		}
		queue = newQ
	}
	return ans
}
