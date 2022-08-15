package leetcode

import (
	"container/heap"
	"math"
)

func KClosest(points [][]int, k int) [][]int {
	h := make(PointHeap, len(points))
	for i, point := range points {
		h[i] = Point{x: point[0], y: point[1]}
	}
	heap.Init(&h)
	ans := make([][]int, k)
	for i := 0; i < k; i++ {
		point := heap.Pop(&h).(Point)
		ans[i] = []int{point.x, point.y}
	}
	return ans
}

type Point struct {
	x int
	y int
}

type PointHeap []Point

func (h PointHeap) Len() int { return len(h) }
func (h PointHeap) Less(i, j int) bool {
	return math.Sqrt(math.Pow(float64(h[i].x), 2.0)+math.Pow(float64(h[i].y), 2.0)) < math.Sqrt(math.Pow(float64(h[j].x), 2.0)+math.Pow(float64(h[j].y), 2.0))
}
func (h PointHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *PointHeap) Push(x interface{}) { *h = append(*h, x.(Point)) }
func (h *PointHeap) Pop() interface{} {
	n := len(*h)
	min := (*h)[n-1]
	*h = (*h)[:n-1]
	return min
}
