package main

func findOrder(numCourses int, prerequisites [][]int) []int {
	graph := make([][]int, numCourses)
	indeg := make([]int, numCourses)

	// Initialize the graph and indeg.
	for _, prereq := range prerequisites {
		from, to := prereq[1], prereq[0]
		graph[from] = append(graph[from], to)
		indeg[to]++
	}

	var res []int

	// Top sort.
	for len(res) < numCourses {
		idx := getZeroIdx(indeg)
		if idx == -1 {
			return []int{}
		}

		res = append(res, idx)

		for _, val := range graph[idx] {
			indeg[val]--
		}
	}
	return res
}

func getZeroIdx(nums []int) int {
	for i, val := range nums {
		if val == 0 {
			nums[i] = -1
			return i
		}
	}
	return -1
}
