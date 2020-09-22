package main

func majorityElement(nums []int) []int {
	var count1, count2 int
	var candidate1, candidate2 *int

	for i, n := range nums {
		if candidate1 != nil && *candidate1 == n {
			count1++
		} else if candidate2 != nil && *candidate2 == n {
			count2++
		} else if count1 == 0 {
			candidate1 = &nums[i]
			count1++
		} else if count2 == 0 {
			candidate2 = &nums[i]
			count2++
		} else {
			count1--
			count2--
		}
	}

	result := []int{}
	count1 = 0
	count2 = 0

	for _, n := range nums {
		if candidate1 != nil && n == *candidate1 {
			count1++
		}
		if candidate2 != nil && n == *candidate2 {
			count2++
		}
	}

	n := len(nums)
	if count1 > n/3 {
		result = append(result, *candidate1)
	}
	if count2 > n/3 {
		result = append(result, *candidate2)
	}

	return result
}
