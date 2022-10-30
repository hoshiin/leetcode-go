package leetcode

func TotalFruit(fruits []int) int {
	left, right := 0, 0
	basket := make(map[int]int)
	for right = 0; right < len(fruits); right++ {
		basket[fruits[right]]++
		if len(basket) > 2 {
			basket[fruits[left]]--
			num := basket[fruits[left]]
			if num == 0 {
				delete(basket, fruits[left])
			}
			left++
		}
	}
	return right - left
}
