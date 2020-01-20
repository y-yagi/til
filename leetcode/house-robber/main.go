package main

func rob(nums []int) int {
	rob, norob := 0, 0
	for _, num := range nums {
		rob, norob = num+norob, max(norob, rob)
	}
	return max(norob, rob)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
