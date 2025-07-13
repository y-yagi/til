package main

import "fmt"

func moveZeroes(nums []int) {
	left := 0
	for right := 0; right < len(nums); right++ {
		if nums[right] != 0 {
			t := nums[left]
			nums[left] = nums[right]
			nums[right] = t
			left++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	fmt.Printf("%v\n", nums)
}
