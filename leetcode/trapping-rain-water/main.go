package main

func trap(height []int) int {
	left := 0
	right := len(height) - 1
	ans := 0
	var leftMax, rightMax int

	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				ans += (leftMax - height[left])
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				ans += (rightMax - height[right])
			}
			right--
		}
	}

	return ans
}
