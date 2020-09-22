package main

func majorityElement(nums []int) []int {
	counter := map[int]int{}
	ans := []int{}
	border := len(nums) / 3

	for i := 0; i < len(nums); i++ {
		counter[nums[i]]++
	}

	for k, v := range counter {
		if v > border {
			ans = append(ans, k)
		}

	}

	return ans
}
