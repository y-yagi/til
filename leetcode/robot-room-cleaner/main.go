package main

import "fmt"

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

func cleanRoom(robot *Robot) {
	visited := map[string]bool{}
	backtrace(0, 0, 0, robot, &visited)
}

func backtrace(row, col, d int, robot *Robot, visited *map[string]bool) {
	directions := [][]int{[]int{1, 0}, []int{-1, 0}, []int{0, 1}, []int{0, -1}}
	current := fmt.Sprintf("%d%d", row, col)
	(*visited)[current] = true
	robot.Clean()

	for i := 0; i < 4; i++ {
		newD := (d + i) % 4
		newRow := row + directions[newD][0]
		newCol := col + directions[newD][1]
		newCurrent := fmt.Sprintf("%d%d", newRow, newCol)

		if _, found := (*visited)[newCurrent]; !found && robot.Move() {
			backtrace(newRow, newCol, newD, robot, visited)
			goBack(robot)
		}
		robot.TurnRight()
	}
}

func goBack(robot *Robot) {
	robot.TurnRight()
	robot.TurnRight()
	robot.Move()
	robot.TurnRight()
	robot.TurnRight()
}
