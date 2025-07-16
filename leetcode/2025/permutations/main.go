package main

func permute(nums []int) [][]int {
	ans := [][]int{}
	backtrack(nums, nil, &ans)
	return ans
}

func backtrack(nums []int, cur []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, append([]int{}, cur...))
		return
	}

	for i := 0; i < len(nums); i++ {
		nums_without_cur_value := append(append([]int{}, nums[:i]...), nums[i+1:]...)
		backtrack(nums_without_cur_value, append(cur, nums[i]), ans)
	}
}
