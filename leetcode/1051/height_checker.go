package leetcode

func HeightChecker(heights []int) int {
	expect := make([]int, len(heights))
	for i := 0; i < len(heights); i++ {
		expect[i] = heights[i]
	}

	for i := 0; i < len(expect); i++ {
		for j := i + 1; j < len(expect); j++ {
			if expect[i] > expect[j] {
				expect[i], expect[j] = expect[j], expect[i]
			}
		}
	}

	count := 0
	for i := 0; i < len(heights); i++ {
		if heights[i] != expect[i] {
			count++
		}
	}
	return count
}

func HeightCheckerWithoutSorting(heights []int) int {
	m := [101]int{}
	for _, h := range heights {
		m[h]++
	}

	i := 0
	count := 0

	for curr := 1; curr <= 100; curr++ {
		for j := m[curr]; j > 0; j-- {
			if heights[i] != curr {
				count++
			}
			i++
		}
	}
	return count
}
