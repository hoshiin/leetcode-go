package leetcode

func ReverseString(s []byte) {
	reverseStringInternal(0, s)
}

func reverseStringInternal(i int, s []byte) {
	if i >= len(s)/2 {
		return
	}
	s[i], s[len(s)-1-i] = s[len(s)-1-i], s[i]
	reverseStringInternal(i+1, s)
}
