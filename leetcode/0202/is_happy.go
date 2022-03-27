package leetcode

func IsHappy(n int) bool {
	seen := make(map[int]struct{})
	for !isSeen(seen, n) && n != 1 {
		seen[n] = struct{}{}
		n = getNext(n)
	}
	return n == 1
}

func getNext(n int) int {
	totalSum := 0
	for n > 0 {
		d := n % 10
		n = n / 10
		totalSum += d * d
	}
	return totalSum
}

func isSeen(seen map[int]struct{}, n int) bool {
	_, ok := seen[n]
	return ok
}

func IsHappyCircleFinding(n int) bool {
	slowRunner := n
	fastRunner := getNext(n)
	for fastRunner != 1 && slowRunner != fastRunner {
		slowRunner = getNext(slowRunner)
		fastRunner = getNext(getNext(fastRunner))
	}
	return fastRunner == 1
}
