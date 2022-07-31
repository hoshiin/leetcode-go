package leetcode

func IsValidSudoku(board [][]byte) bool {
	n := 9
	var rows, cols, boxes []map[byte]struct{}
	for i := 0; i < 9; i++ {
		rows = append(rows, make(map[byte]struct{}))
		cols = append(cols, make(map[byte]struct{}))
		boxes = append(boxes, make(map[byte]struct{}))
	}
	for r := 0; r < n; r++ {
		for c := 0; c < n; c++ {
			val := board[r][c]
			if val == '.' {
				continue
			}
			if _, ok := rows[r][val]; ok {
				return false
			}
			rows[r][val] = struct{}{}
			if _, ok := cols[c][val]; ok {
				return false
			}
			cols[c][val] = struct{}{}
			index := (r/3)*3 + c/3
			if _, ok := boxes[index][val]; ok {
				return false
			}
			boxes[index][val] = struct{}{}
		}
	}
	return true
}
