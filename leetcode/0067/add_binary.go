package leetcode

func AddBinary(a string, b string) string {
	i, j := len(a), len(b)
	ans := ""
	carry := false
	for i > 0 || j > 0 {
		i--
		j--
		count := 0
		if i >= 0 && a[i] == '1' {
			count++
		}
		if j >= 0 && b[j] == '1' {
			count++
		}

		if carry {
			count++
			carry = false
		}
		switch count {
		case 0:
			ans = "0" + ans
		case 1:
			ans = "1" + ans
		case 2:
			ans = "0" + ans
			carry = true
		default:
			ans = "1" + ans
			carry = true
		}
	}
	if carry {
		ans = "1" + ans
	}
	return ans
}
