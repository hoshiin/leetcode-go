package leetcode

func CanPlaceFlowers(flowerbed []int, n int) bool {
	count := n
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			emptyLeft := i == 0 || flowerbed[i-1] == 0
			emptyRight := i == len(flowerbed)-1 || flowerbed[i+1] == 0
			if emptyLeft && emptyRight {
				flowerbed[i] = 1
				count--
			}
		}
		if count == 0 {
			return true
		}
	}
	return count <= 0
}
