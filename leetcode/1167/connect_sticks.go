package leetcode

import "container/heap"

func ConnectSticks(sticks []int) int {
	s := IntHeap(sticks)
	heap.Init(&s)
	sum := 0
	for s.Len() > 1 {
		i, j := heap.Pop(&s).(int), heap.Pop(&s).(int)
		sum += i + j
		heap.Push(&s, i+j)
	}
	return sum
}

type IntHeap []int

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() interface{} {
	n := len(*h)
	pop := (*h)[n-1]
	*h = (*h)[:n-1]
	return pop
}
