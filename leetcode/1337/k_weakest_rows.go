package leetcode

import "container/heap"

func KWeakestRows(mat [][]int, k int) []int {
	rows := make(Rows, len(mat))
	for i, row := range mat {
		count := 0
		for _, r := range row {
			if r == 1 {
				count++
			}
		}
		rows[i] = Row{
			i:        i,
			soldiers: count,
			row:      row,
		}
	}
	heap.Init(&rows)
	ans := []int{}
	for len(ans) < k {
		ans = append(ans, heap.Pop(&rows).(Row).i)
	}
	return ans
}

type Row struct {
	i        int
	soldiers int
	row      []int
}

type Rows []Row

func (r Rows) Len() int      { return len(r) }
func (r Rows) Swap(i, j int) { r[i], r[j] = r[j], r[i] }
func (r Rows) Less(i, j int) bool {
	if r[i].soldiers != r[j].soldiers {
		return r[i].soldiers < r[j].soldiers
	}
	for n, row := range r[i].row {
		if row != r[j].row[n] {
			return row != 1
		}
	}
	return r[i].i < r[j].i
}
func (r *Rows) Push(x interface{}) {
	*r = append(*r, x.(Row))
}
func (r *Rows) Pop() interface{} {
	n := len(*r)
	min := (*r)[n-1]
	*r = (*r)[:n-1]
	return min
}
