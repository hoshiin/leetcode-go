package leetcode

func CountVowelPermutation(n int) int {
	aVowelPermutationCount := make([]uint64, n)
	eVowelPermutationCount := make([]uint64, n)
	iVowelPermutationCount := make([]uint64, n)
	oVowelPermutationCount := make([]uint64, n)
	uVowelPermutationCount := make([]uint64, n)

	aVowelPermutationCount[0] = 1
	eVowelPermutationCount[0] = 1
	iVowelPermutationCount[0] = 1
	oVowelPermutationCount[0] = 1
	uVowelPermutationCount[0] = 1

	mod := uint64(1000000007)

	for i := 1; i < n; i++ {
		aVowelPermutationCount[i] = (eVowelPermutationCount[i-1] + iVowelPermutationCount[i-1] + uVowelPermutationCount[i-1]) % mod
		eVowelPermutationCount[i] = (aVowelPermutationCount[i-1] + iVowelPermutationCount[i-1]) % mod
		iVowelPermutationCount[i] = (eVowelPermutationCount[i-1] + oVowelPermutationCount[i-1]) % mod
		oVowelPermutationCount[i] = (iVowelPermutationCount[i-1]) % mod
		uVowelPermutationCount[i] = (iVowelPermutationCount[i-1] + oVowelPermutationCount[i-1]) % mod
	}
	result := uint64(0)
	result = (aVowelPermutationCount[n-1] + eVowelPermutationCount[n-1] + iVowelPermutationCount[n-1] + oVowelPermutationCount[n-1] + uVowelPermutationCount[n-1]) % mod
	return int(result)
}
