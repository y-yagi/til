package main

func intersect(nums1 []int, nums2 []int) []int {
	result := []int{}
	dict := map[int]int{}

	for i := 0; i < len(nums1); i++ {
		dict[nums1[i]]++
	}

	for i := 0; i < len(nums2); i++ {
		if _, found := dict[nums2[i]]; found {
			if dict[nums2[i]] > 0 {
				result = append(result, nums2[i])
				dict[nums2[i]]--
			}
		}
	}

	return result
}
