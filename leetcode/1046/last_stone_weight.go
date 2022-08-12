package leetcode

import "container/heap"

func LastStoneWeight(stones []int) int {
	s := IntHeap(stones)
	heap.Init(&s)
	for len(s) > 1 {
		y, x := heap.Pop(&s).(int), heap.Pop(&s).(int)
		if x != y {
			heap.Push(&s, y-x)
		}
	}
	if len(s) == 0 {
		return 0
	}
	return s[0]
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] > h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	top := (*h)[n-1]
	*h = (*h)[:n-1]
	return top
}
