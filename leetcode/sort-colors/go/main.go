package main

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	curr := 0

	var t int

	for curr <= r {
		if nums[curr] == 0 {
			t = nums[l]
			nums[l] = nums[curr]
			nums[curr] = t
			l++
			curr++
		} else if nums[curr] == 2 {
			t = nums[curr]
			nums[curr] = nums[r]
			nums[r] = t
			r--
		} else {
			curr++
		}
	}
}
