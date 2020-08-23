package main

func hasPath(maze [][]int, start []int, destination []int) bool {
	visited := [][]bool{}
	for i := 0; i < len(maze); i++ {
		visited = append(visited, make([]bool, len(maze[0])))
	}

	return dfs(maze, start, destination, visited)
}

func dfs(maze [][]int, start []int, destination []int, visited [][]bool) bool {
	if visited[start[0]][start[1]] {
		return false
	}

	if start[0] == destination[0] && start[1] == destination[1] {
		return true
	}

	visited[start[0]][start[1]] = true

	r := start[1] + 1
	l := start[1] - 1
	u := start[0] - 1
	d := start[0] + 1

	// right
	for r < len(maze[0]) && maze[start[0]][r] == 0 {
		r++
	}
	if dfs(maze, []int{start[0], r - 1}, destination, visited) {
		return true
	}

	// left
	for l >= 0 && maze[start[0]][l] == 0 {
		l--
	}
	if dfs(maze, []int{start[0], l + 1}, destination, visited) {
		return true
	}

	// up
	for u >= 0 && maze[u][start[1]] == 0 {
		u--
	}
	if dfs(maze, []int{u + 1, start[1]}, destination, visited) {
		return true
	}

	// down
	for d < len(maze) && maze[d][start[1]] == 0 {
		d++
	}
	if dfs(maze, []int{d - 1, start[1]}, destination, visited) {
		return true
	}
	return false
}
