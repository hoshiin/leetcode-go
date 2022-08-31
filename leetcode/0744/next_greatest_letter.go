package leetcode

func NextGreatestLetter(letters []byte, target byte) byte {
	low, high := 0, len(letters)
	for low < high {
		mid := low + (high-low)/2
		if letters[mid] <= target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return letters[low%len(letters)]
}
