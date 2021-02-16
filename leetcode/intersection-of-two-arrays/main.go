package main

func intersection(nums1 []int, nums2 []int) []int {
	ans := []int{}
	dict1 := map[int]bool{}
	dict2 := map[int]bool{}

	for i := 0; i < len(nums1); i++ {
		dict1[nums1[i]] = true
	}

	for i := 0; i < len(nums2); i++ {
		if _, found := dict1[nums2[i]]; found {
			if _, found := dict2[nums2[i]]; !found {
				ans = append(ans, nums2[i])
			}
			dict2[nums2[i]] = true
		}
	}

	return ans
}
