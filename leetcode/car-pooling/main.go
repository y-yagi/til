package main

func carPooling(trips [][]int, capacity int) bool {
	station := make([]int, 1001)
	for _, trip := range trips {
		station[trip[1]] += trip[0]
		station[trip[2]] -= trip[0]
	}
	passengers := 0
	for _, v := range station {
		passengers += v
		if passengers > capacity {
			return false
		}
	}
	return true
}
