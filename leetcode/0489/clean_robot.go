package leetcode

// This is the robot's control interface.
// You should not implement it, or speculate about its implementation
type Robot struct {
}

// Returns true if the cell in front is open and robot moves into the cell.
// Returns false if the cell in front is blocked and robot stays in the current cell.
func (robot *Robot) Move() bool { return true }

// Robot will stay in the same cell after calling TurnLeft/TurnRight.
// Each turn will be 90 degrees.
func (robot *Robot) TurnLeft()  {}
func (robot *Robot) TurnRight() {}

// Clean the current cell.
func (robot *Robot) Clean() {}

func CleanRoom(robot *Robot) {
	backtrack(robot, 0, 0, 0, map[pair]struct{}{})
}

type pair struct {
	row int
	col int
}

// going clockwise : 0: 'up', 1: 'right', 2: 'down', 3: 'left'
var directions = [][]int{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}

func backtrack(robot *Robot, row, col, d int, visited map[pair]struct{}) {
	visited[pair{row: row, col: col}] = struct{}{}
	robot.Clean()
	// going clockwise : 0: 'up', 1: 'right', 2: 'down', 3: 'left'
	for i := 0; i < 4; i++ {
		newD := (d + i) % 4
		newRow := row + directions[newD][0]
		newCol := col + directions[newD][1]

		if _, ok := visited[pair{row: newRow, col: newCol}]; !ok && robot.Move() {
			backtrack(robot, newRow, newCol, newD, visited)
			goBack(robot)
		}
		robot.TurnRight()
	}
}
