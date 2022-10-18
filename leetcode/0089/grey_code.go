package leetcode

func GrayCode(n int) []int {
	result := []int{0}
	isPresent := map[int]struct{}{0: {}}
	grayCodeHelper(&result, n, isPresent)
	return result
}

func grayCodeHelper(result *[]int, n int, isPresent map[int]struct{}) bool {
	l := len(*result)
	if l == (1 << n) {
		return true
	}
	current := (*result)[l-1]
	for i := 0; i < n; i++ {
		next := current ^ (1 << i)
		if _, ok := isPresent[next]; !ok {
			isPresent[next] = struct{}{}
			*result = append(*result, next)
			if grayCodeHelper(result, n, isPresent) {
				return true
			}
			delete(isPresent, next)
			*result = (*result)[:l-1]
		}
	}
	return false
}
