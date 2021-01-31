package main

func subsets(nums []int) [][]int {
	ans := [][]int{}
	backtrace(nums, 0, []int{}, &ans)
	return ans
}

func backtrace(nums []int, index int, curr []int, ans *[][]int) {
	dup := make([]int, len(curr))
	copy(dup, curr)
	*ans = append(*ans, dup)

	for ; index < len(nums); index++ {
		curr = append(curr, nums[index])
		backtrace(nums, index+1, curr, ans)
		curr = curr[:len(curr)-1]
	}
}
