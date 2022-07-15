package main

func canVisitAllRooms(rooms [][]int) bool {
	if len(rooms) == 0 {
		return true
	}

	visited := make([]bool, len(rooms))
	visited[0] = true
	queue := rooms[0]

	for len(queue) != 0 {
		newKeys := []int{}
		for _, k := range queue {
			if !visited[k] {
				visited[k] = true
				newKeys = append(newKeys, rooms[k]...)
			}
		}
		queue = newKeys
	}

	for _, v := range visited {
		if v == false {
			return false
		}
	}

	return true
}
