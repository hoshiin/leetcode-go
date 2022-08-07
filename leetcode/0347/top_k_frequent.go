package leetcode

import "container/heap"

func TopKFrequent(nums []int, k int) []int {
	seen := map[int]int{}
	for _, num := range nums {
		seen[num]++
	}
	q := PriorityQueue{}
	for key, count := range seen {
		heap.Push(&q, &Item{key: key, count: count})
	}
	ans := []int{}
	for len(ans) < k {
		item := heap.Pop(&q).(*Item)
		ans = append(ans, item.key)
	}
	return ans
}

type Item struct {
	key   int
	count int
}

type PriorityQueue []*Item

func (q PriorityQueue) Len() int { return len(q) }

func (q PriorityQueue) Less(i, j int) bool { return q[i].count > q[j].count }

func (q PriorityQueue) Swap(i, j int) { q[i], q[j] = q[j], q[i] }

func (q *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*q = append(*q, item)
}

func (q *PriorityQueue) Pop() interface{} {
	n := len(*q)
	item := (*q)[n-1]
	*q = (*q)[:n-1]
	return item
}
