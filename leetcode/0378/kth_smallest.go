package leetcode

import (
	"container/heap"
	"sort"
)

func KthSmallest(matrix [][]int, k int) int {
	arr := []int{}
	for _, row := range matrix {
		arr = append(arr, row...)
	}
	sort.Ints(arr)
	return arr[k-1]
}

func KthSmallestHeap(matrix [][]int, k int) int {
	arr := pq([]int{})
	heap.Init(&arr)
	for _, row := range matrix {
		for _, n := range row {
			if arr.Len() < k {
				heap.Push(&arr, n)
			} else if n < arr[0] {
				heap.Pop(&arr)
				heap.Push(&arr, n)
			} else {
				break
			}
		}
	}
	return heap.Pop(&arr).(int)
}

type pq []int

func (p pq) Len() int            { return len(p) }
func (p pq) Less(i, j int) bool  { return p[i] > p[j] }
func (p pq) Swap(i, j int)       { p[i], p[j] = p[j], p[i] }
func (p *pq) Push(x interface{}) { *p = append(*p, x.(int)) }
func (p *pq) Pop() interface{} {
	n := len(*p)
	min := (*p)[n-1]
	*p = (*p)[:n-1]
	return min
}
