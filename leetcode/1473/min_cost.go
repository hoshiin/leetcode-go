package leetcode

func MinCost(houses []int, cost [][]int, m int, n int, target int) int {
	// Maximum cost possible plus 1
	maxCost := 1000001
	prevMemo := make([][]int, target+1)
	// Initialize prevMemo array
	for i := range prevMemo {
		memo := make([]int, n)
		for j := range memo {
			memo[j] = maxCost
		}
		prevMemo[i] = memo
	}
	// Initialize for house 0, neighborhood will be 1
	for color := 1; color <= n; color++ {
		if houses[0] == color {
			// If the house has same color, no cost
			prevMemo[1][color-1] = 0
		} else if houses[0] == 0 {
			// If the house is not painted, assign the corresponding
			prevMemo[1][color-1] = cost[0][color-1]
		}
	}
	for house := 1; house < m; house++ {
		memo := make([][]int, target+1)
		for i := range memo {
			m := make([]int, n)
			// Initialize memo array
			for j := range m {
				m[j] = maxCost
			}
			memo[i] = m
		}
		for neighborhoods := 1; neighborhoods <= min(target, house+1); neighborhoods++ {
			for color := 1; color <= n; color++ {
				if houses[house] != 0 && color != houses[house] {
					// Cannot be painted with different color
					continue
				}
				currCost := maxCost
				// Iterate over all the possible color for previous house
				for prevColor := 1; prevColor <= n; prevColor++ {
					if prevColor != color {
						// Decrement the neighborhood as adjacent houses has different color
						currCost = min(currCost, prevMemo[neighborhoods-1][prevColor-1])
					} else {
						// Previous house has the same color, no change in neighborhood count
						currCost = min(currCost, prevMemo[neighborhoods][color-1])
					}
				}
				// If the house is already painted cost to paint is 0
				costToPaint := 0
				if houses[house] == 0 {
					costToPaint = cost[house][color-1]
				}
				// Update the table to have the current house results
				memo[neighborhoods][color-1] = currCost + costToPaint
			}
		}
		prevMemo = memo
	}

	minCost := maxCost
	// Find the minimum cost with m houses and target neighborhoods
	// By comparing cost for different color for the last house
	for color := 1; color <= n; color++ {
		minCost = min(minCost, prevMemo[target][color-1])
	}
	// Return -1 if the answer is MAX_COST as it implies no answer possible
	if minCost == maxCost {
		return -1
	}
	return minCost
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
