package leetcode

import "sort"

func UniqueOccurrences(arr []int) bool {
	occur := make(map[int]int)
	for _, n := range arr {
		occur[n]++
	}
	unique := make(map[int]struct{})

	for _, o := range occur {
		if _, ok := unique[o]; ok {
			return false
		}
		unique[o] = struct{}{}
	}
	return true
}

func UniqueOccurrencesUseSort(arr []int) bool {
	occur := make(map[int]int)
	for _, n := range arr {
		occur[n]++
	}
	unique := make([]int, 0, len(arr))
	for _, o := range occur {
		unique = append(unique, o)
	}
	sort.Slice(unique, func(i, j int) bool { return unique[i] < unique[j] })
	for i := 1; i < len(unique); i++ {
		if unique[i-1] == unique[i] {
			return false
		}
	}
	return true
}
