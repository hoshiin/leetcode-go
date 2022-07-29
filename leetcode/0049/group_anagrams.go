package leetcode

import "sort"

func GroupAnagrams(strs []string) [][]string {
	m := map[string][]string{}
	for _, str := range strs {
		r := []byte(str)
		sort.Slice(r, func(i, j int) bool { return r[i] < r[j] })
		key := string(r)
		if arr, ok := m[key]; !ok {
			m[key] = []string{str}
		} else {
			m[key] = append(arr, str)
		}
	}
	ans := make([][]string, len(m))
	i := 0
	for _, arr := range m {
		ans[i] = arr
		i++
	}
	return ans
}
