package main

func largestRectangleArea(heights []int) int {
	return calucareArea(heights, 0, len(heights)-1)
}

func calucareArea(heights []int, start, end int) int {
	if start > end {
		return 0
	}

	minindex := start
	for i := start; i <= end; i++ {
		if heights[minindex] > heights[i] {
			minindex = i
		}
	}
	return max(heights[minindex]*(end-start+1), max(calucareArea(heights, start, minindex-1), calucareArea(heights, minindex+1, end)))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
