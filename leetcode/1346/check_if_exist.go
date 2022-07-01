package leetcode

func CheckIfExist(arr []int) bool {
	mp := make(map[float64]bool)
	for i := 0; i < len(arr); i++ {
		v := float64(arr[i])
		if mp[v*2.0] || mp[v/2.0] {
			return true
		}
		mp[v] = true
	}
	return false
}

func CheckIfExistBruteForce(arr []int) bool {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if i != j && arr[i]*2 == arr[j] {
				return true
			}
		}
	}
	return false
}
