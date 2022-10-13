package leetcode

func HammingWeight(num uint32) int {
	bits := 0
	mask := 1
	for i := 0; i < 32; i++ {
		if (num & uint32(mask)) != 0 {
			bits++
		}
		mask <<= 1
	}
	return bits
}
