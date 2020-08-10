package main

func moveZeroes(nums []int) {
	i := 0
	for j := 0; j < len(nums); j++ {
		if nums[j] != 0 {
			t := nums[i]
			nums[i] = nums[j]
			nums[j] = t
			i++
		}
	}
}
