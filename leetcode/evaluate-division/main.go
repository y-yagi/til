package main

type point struct {
	name string
	dis  float64
}

func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	a := equations
	graph := make(map[string]map[string]float64)
	visited := make(map[string]bool)
	for i, v := range a {
		if _, ok := graph[v[0]]; !ok {
			graph[v[0]] = make(map[string]float64)
		}
		graph[v[0]][v[1]] = values[i]
		if _, ok := graph[v[1]]; !ok {
			graph[v[1]] = make(map[string]float64)
		}
		graph[v[1]][v[0]] = 1 / values[i]

		visited[v[0]] = false
		visited[v[1]] = false
	}

	var res []float64
	for _, v := range queries {
		_, ok1 := graph[v[0]]
		_, ok2 := graph[v[1]]
		if !ok1 || !ok2 {
			res = append(res, -1.0)
			continue
		}

		if v[0] == v[1] {
			res = append(res, 1.0)
			continue
		}

		for k, _ := range visited {
			visited[k] = false
		}

		res = append(res, dijkstra(graph, visited, v[0], v[1]))
	}
	return res
}

func dijkstra(graph map[string]map[string]float64, visited map[string]bool, start string, end string) float64 {
	visited[start] = true
	queue := []point{point{start, 1.0}}

	for len(queue) > 0 {
		l := len(queue)
		for i := 0; i < l; i++ {
			v := queue[i]
			for v2 := range graph[v.name] {
				d := graph[v.name][v2]
				if v2 == end {
					return d * v.dis
				}
				if visited[v2] {
					continue
				}
				queue = append(queue, point{v2, d * v.dis})
				visited[v2] = true
			}
		}
		queue = queue[l:]
	}
	return -1.0
}
