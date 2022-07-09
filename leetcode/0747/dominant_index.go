package leetcode

func DominantIndex(nums []int) int {
	largest := -1
	largestIndex := -1
	for i, num := range nums {
		if largest < num {
			largest = num
			largestIndex = i
		}
	}
	for _, num := range nums {
		if num != largest && num*2 > largest {
			return -1
		}
	}
	return largestIndex
}
