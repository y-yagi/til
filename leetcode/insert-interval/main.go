package main

func insert(intervals [][]int, newInterval []int) [][]int {
	l := len(intervals)
	for i := 0; i < l; i++ {
		if intervals[i][0] > newInterval[0] {
			intervals = append(intervals[:i], append([][]int{newInterval}, intervals[i:]...)...)
			break
		}
	}
	if len(intervals) == l {
		intervals = append(intervals, newInterval)
	}
	return merge(intervals)
}

func merge(intervals [][]int) [][]int {
	answer := make([][]int, 0)
	for i := 0; i < len(intervals); i++ {
		if len(answer) == 0 || answer[len(answer)-1][1] < intervals[i][0] {
			answer = append(answer, intervals[i])
		} else {
			answer[len(answer)-1][1] = max(answer[len(answer)-1][1], intervals[i][1])
		}
	}
	return answer
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
