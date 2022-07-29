package main

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	answer := []int{}
	dict := map[int]int{}

	for i := 0; i < len(nums2); i++ {
		dict[nums2[i]] = i
	}

	for i := 0; i < len(nums1); i++ {
		j, _ := dict[nums1[i]]

		for ; j < len(nums2); j++ {
			if nums1[i] < nums2[j] {
				answer = append(answer, nums2[j])
				break
			}
		}

		if len(answer) <= i {
			answer = append(answer, -1)
		}
	}

	return answer
}
