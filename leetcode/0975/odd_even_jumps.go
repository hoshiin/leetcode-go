package leetcode

func OddEvenJumps(arr []int) int {
	odd, even := make([]bool, len(arr)), make([]bool, len(arr))
	odd[len(arr)-1], even[len(arr)-1] = true, true

	for i := len(arr) - 2; i >= 0; i-- {
		jodd, jeven := -1, -1
		for j := i + 1; j < len(arr); j++ {
			if arr[j] >= arr[i] {
				if jodd == -1 || arr[j] < arr[jodd] || arr[j] == arr[jodd] && j < jodd {
					jodd = j
					odd[i] = even[j]
				}
			}
			if arr[j] <= arr[i] {
				if jeven == -1 || arr[j] > arr[jeven] || arr[j] == arr[jeven] && j < jeven {
					jeven = j
					even[i] = odd[j]
				}
			}
		}
	}
	res := 0
	for _, v := range odd {
		if v {
			res++
		}
	}
	return res
}
