package leetcode

func DuplicateZeros(arr []int) {
	count := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			count++
		}
	}
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == 0 {
			if count+i < len(arr) {
				arr[count+i] = 0
			}
			count--
		}
		if i+count < len(arr) {
			arr[count+i] = arr[i]
		}
	}
}
