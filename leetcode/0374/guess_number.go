package leetcode

/**
 * Forward declaration of guess API.
 * @param  num   your guess
 * @return 	     -1 if num is higher than the picked number
 *			      1 if num is lower than the picked number
 *               otherwise return 0
 * func guess(num int) int;
 */
var pick int

func GuessNumber(n int) int {
	low, high := 1, n
	for low <= high {
		mid := low + (high-low)/2
		res := guess(mid)
		if res == 0 {
			return mid
		} else if res < 0 {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return -1
}

func guess(n int) int {
	if pick == n {
		return 0
	} else if pick < n {
		return 1
	}
	return -1
}
