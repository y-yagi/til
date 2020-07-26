package main

func findMin(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}

	left := 0
	right := len(nums) - 1

	for right > left {
		mid := left + (right-left)/2
		if nums[mid] < nums[right] {
			right = mid
		} else if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			right -= 1
		}
	}

	return nums[left]
}
