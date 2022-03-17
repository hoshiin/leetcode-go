package leetcode

var isBadVersion = func(v int) bool {
	return v == version
}

var version int

func FirstBadVersion(n int) int {
	left, right := 0, n
	for left < right {
		mid := (right-left)/2 + left
		if isBadVersion(mid) {
			if mid-1 == 0 || !isBadVersion(mid-1) {
				return mid
			}
			right = mid
			continue
		}
		left = mid + 1
	}
	return left
}
