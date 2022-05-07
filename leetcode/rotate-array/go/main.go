package main

func rotate(nums []int, k int) {
	l := len(nums)
	if l != 1 {
		clone := nums[l-(k%l):]
		latter := nums[:l-(k%l)]
		clone = append(clone, latter...)
		copy(nums, clone)
	}
}
