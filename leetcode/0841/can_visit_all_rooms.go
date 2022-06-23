package leetcode

func CanVisitAllRooms(rooms [][]int) bool {
	seen := make([]bool, len(rooms))
	seen[0] = true
	stack := []int{0}
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, key := range rooms[node] {
			if !seen[key] {
				seen[key] = true
				stack = append(stack, key)
			}
		}
	}

	for _, room := range seen {
		if !room {
			return false
		}
	}
	return true
}
