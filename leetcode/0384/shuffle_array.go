package leetcode

import "math/rand"

type Solution struct {
	arr []int
}

func Constructor(nums []int) Solution {
	return Solution{
		arr: nums,
	}
}

func (this *Solution) Reset() []int {
	return this.arr
}

func (this *Solution) Shuffle() []int {
	nums := make([]int, len(this.arr))
	copy(nums, this.arr)
	for i := 0; i < len(nums); i++ {
		index := rand.Intn(len(nums))
		nums[i], nums[index] = nums[index], nums[i]
	}
	return nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
