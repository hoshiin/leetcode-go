package leetcode

import "container/heap"

type KthLargest struct {
	k    int
	nums *MinHeap
}

func Constructor(k int, nums []int) KthLargest {
	n := MinHeap(nums)
	heap.Init(&n)
	for len(n) > k {
		heap.Pop(&n)
	}
	return KthLargest{k: k, nums: &n}
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this.nums, val)
	for len(*this.nums) > this.k {
		heap.Pop(this.nums)
	}
	return (*this.nums)[0]
}

type MinHeap []int

func (h MinHeap) Len() int            { return len(h) }
func (h MinHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MinHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MinHeap) Pop() interface{} {
	old := *h
	n := len(old)
	min := old[n-1]
	*h = old[:n-1]
	return min
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
