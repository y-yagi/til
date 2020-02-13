package main

func intersection(nums1 []int, nums2 []int) []int {
	result := []int{}
	dict := map[int]bool{}
	answer := map[int]bool{}

	for i := 0; i < len(nums1); i++ {
		dict[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if _, found := dict[nums2[i]]; found {
			if _, found := answer[nums2[i]]; !found {
				result = append(result, nums2[i])
				answer[nums2[i]] = true
			}
		}
	}

	return result
}
