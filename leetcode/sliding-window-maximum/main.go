package main

func maxSlidingWindow(nums []int, k int) []int {
	var res []int
	var window []int
	for index := 0; index < len(nums); index++ {
		if index >= k && window[0] <= index-k {
			window = window[1:]
		}
		for len(window) > 0 && nums[window[len(window)-1]] < nums[index] {
			window = window[:len(window)-1]
		}
		window = append(window, index)
		if index >= k-1 {
			res = append(res, nums[window[0]])
		}
	}
	return res
}
