package main

func intersection(nums1 []int, nums2 []int) []int {
	numMap := make(map[int]int)
	for _, num := range nums1 {
		numMap[num]++
	}

	var result []int
	for _, num := range nums2 {
		if count, exists := numMap[num]; exists && count > 0 {
			result = append(result, num)
			numMap[num] = 0
		}
	}

	return result
}
