package main

import "fmt"

func nextPermutation(nums []int) {
	i := len(nums) - 2

	for ; i >= 0 && nums[i+1] <= nums[i]; i-- {
	}

	if i >= 0 {
		j := len(nums) - 1
		for ; j >= 0 && nums[j] <= nums[i]; j-- {
		}
		swap(nums, i, j)
	}
	reverse(nums, i+1)
}

func swap(nums []int, i, j int) {
	t := nums[i]
	nums[i] = nums[j]
	nums[j] = t
}

func reverse(nums []int, start int) {
	i := start
	j := len(nums) - 1

	for i < j {
		swap(nums, i, j)
		i++
		j--
	}
}

func main() {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	fmt.Printf("%v\n", nums)
}
