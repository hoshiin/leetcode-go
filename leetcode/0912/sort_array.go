package leetcode

func SortArray(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	first := SortArray(nums[:len(nums)/2])
	second := SortArray(nums[len(nums)/2:])
	return merge(first, second)
}

func merge(a, b []int) []int {
	ans := []int{}
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] < b[j] {
			ans = append(ans, a[i])
			i++
		} else {
			ans = append(ans, b[j])
			j++
		}
	}
	for ; i < len(a); i++ {
		ans = append(ans, a[i])
	}
	for ; j < len(b); j++ {
		ans = append(ans, b[j])
	}
	return ans
}

func SortArrayQuickSort(nums []int) []int {
	quickSort(nums, 0, len(nums)-1)
	return nums
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, low, pivot-1)
		quickSort(arr, pivot+1, high)
	}

}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return i
}
