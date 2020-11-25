package main

import "sort"

func fourSum(nums []int, target int) [][]int {
	if len(nums) < 4 {
		return nil
	}
	sort.Ints(nums)
	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; {
		for j := i + 1; j < len(nums)-2; {
			l, r := j+1, len(nums)-1
			for l < r {
				if sum := nums[i] + nums[j] + nums[l] + nums[r]; sum == target {
					res = append(res, []int{nums[i], nums[j], nums[l], nums[r]})
					// Skip duplicates
					for l < r && nums[l] == nums[l+1] {
						l++
					}
					l++
					for l < r && nums[r] == nums[r-1] {
						r--
					}
					r--
				} else if sum < target {
					l++
				} else {
					r--
				}
			}
			// Skip duplicates.
			for j < len(nums)-2 && nums[j] == nums[j+1] {
				j++
			}
			j++
		}
		// Skip duplicates.
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
		i++
	}
	return res
}
