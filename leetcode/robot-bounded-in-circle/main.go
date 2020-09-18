package main

func isRobotBounded(instructions string) bool {
	directories := [][]int{[]int{0, 1}, []int{1, 0}, []int{0, -1}, []int{-1, 0}}
	var x, y, idx int

	for _, i := range instructions {
		if i == 'L' {
			idx = (idx + 3) % 4
		} else if i == 'R' {
			idx = (idx + 1) % 4
		} else {
			x += directories[idx][0]
			y += directories[idx][1]
		}
	}

	// after one circle: robot returns into initial position
	// or robot doesn't face norht.
	return (x == 0 && y == 0) || idx != 0
}
