package main

func maxDistance(arrays [][]int) int {
	ans := 0
	minVal := arrays[0][0]
	maxVal := arrays[0][len(arrays[0])-1]
	for i := 1; i < len(arrays); i++ {
		l := len(arrays[i]) - 1
		ans = max(ans, max(abs(arrays[i][l]-minVal), abs(maxVal-arrays[i][0])))
		minVal = min(minVal, arrays[i][0])
		maxVal = max(maxVal, arrays[i][l])
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}

	return x
}
