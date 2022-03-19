package leetcode

func MoveZeroes(nums []int) {
	numZeros := 0
	for _, n := range nums {
		if n == 0 {
			numZeros++
		}
	}
	count := 0
	for _, n := range nums {
		if n != 0 {
			nums[count] = n
			count++
		}
	}

	for numZeros > 0 {
		nums[len(nums)-numZeros] = 0
		numZeros--
	}
}

func MoveZeroesSpaceOptimal(nums []int) {
	lastNonZeroFoundAt := 0
	for _, n := range nums {
		if n != 0 {
			nums[lastNonZeroFoundAt] = n
			lastNonZeroFoundAt++
		}
	}
	for i := lastNonZeroFoundAt; i < len(nums); i++ {
		nums[i] = 0
	}
}

func MoveZeroesSwap(nums []int) {
	lastNonZeroFoundAt := 0
	for cur := 0; cur < len(nums); cur++ {
		if nums[cur] != 0 {
			nums[lastNonZeroFoundAt], nums[cur] = nums[cur], nums[lastNonZeroFoundAt]
			lastNonZeroFoundAt++
		}
	}
}
