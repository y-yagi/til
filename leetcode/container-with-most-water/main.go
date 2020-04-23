package main

func maxArea(height []int) int {
	ans := 0
	l := 0
	r := len(height) - 1

	for r > l {
		area := 0
		if height[l] > height[r] {
			area = height[r] * (r - l)
			r--
		} else {
			area = height[l] * (r - l)
			l++
		}
		ans = max(area, ans)
	}

	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
