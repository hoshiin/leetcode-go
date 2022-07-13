package leetcode

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	length := len(needle)
	for i := 0; i < len(haystack)-len(needle); i++ {
		if haystack[i:i+length] == needle {
			return i
		}
	}
	return -1
}
