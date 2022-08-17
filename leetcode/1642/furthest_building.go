package leetcode

import "container/heap"

func FurthestBuilding(heights []int, bricks int, ladders int) int {
	pq := PriorityQueue([]int{})
	heap.Init(&pq)
	for i := 0; i < len(heights)-1; i++ {
		climb := heights[i+1] - heights[i]
		if climb <= 0 {
			continue
		}
		heap.Push(&pq, climb)
		if pq.Len() <= ladders {
			continue
		}
		bricks -= heap.Pop(&pq).(int)
		if bricks < 0 {
			return i
		}
	}
	return len(heights) - 1
}

type PriorityQueue []int

func (q PriorityQueue) Len() int            { return len(q) }
func (q PriorityQueue) Less(i, j int) bool  { return q[i] < q[j] }
func (q PriorityQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *PriorityQueue) Push(x interface{}) { *q = append(*q, x.(int)) }
func (q *PriorityQueue) Pop() interface{} {
	n := len(*q)
	p := (*q)[n-1]
	*q = (*q)[:n-1]
	return p
}
