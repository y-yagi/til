package main

func subarraySum(nums []int, k int) int {
	var count, sum int
	// Keep count of occurrences of sum
	dict := map[int]int{}
	dict[0] = 1

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if v, found := dict[sum-k]; found {
			count += v
		}

		if v, found := dict[sum]; found {
			dict[sum] = v + 1
		} else {
			dict[sum] = 1
		}
	}

	return count
}
