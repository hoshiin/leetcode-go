package leetcode

func LicenseKeyFormatting(s string, k int) string {
	arr := []byte{}
	count := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '-' {
			continue
		}
		if count == k {
			arr = append(arr, '-')
			count = 0
		}
		if 'a' <= s[i] && s[i] <= 'z' {
			arr = append(arr, byte(s[i]-('a'-'A')))
		} else {
			arr = append(arr, s[i])
		}
		count++
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr)
}
