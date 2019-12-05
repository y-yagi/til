package main

import "fmt"

func main() {
	fmt.Printf("%v\n", twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Printf("%v\n", twoSum([]int{15, 1, 11, 0, 8}, 9))
}

func twoSum(nums []int, target int) []int {
	var fvalue, findex int
	var answer []int

	for {
		if len(nums) == 0 {
			break
		}

		fvalue, nums = nums[0], nums[1:]

		for i, num := range nums {
			if fvalue+num == target {
				answer = append(answer, findex, i+findex+1)
				break
			}
		}
		findex++
	}

	return answer
}
