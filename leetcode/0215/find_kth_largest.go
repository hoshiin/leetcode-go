package leetcode

import "container/heap"

func FindKthLargest(nums []int, k int) int {
	h := MaxHeap(nums)
	heap.Init(&h)
	for i := 0; len(nums) > 0; i++ {
		n := heap.Pop(&h)
		if i == k-1 {
			return n.(int)
		}
	}
	return -1
}

type MaxHeap []int

func (h MaxHeap) Len() int            { return len(h) }
func (h MaxHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	max := old[n-1]
	*h = old[:n-1]
	return max
}
