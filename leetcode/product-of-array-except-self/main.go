package main

func productExceptSelf(nums []int) []int {
	l := len(nums)
	L := make([]int, l)
	R := make([]int, l)
	ans := make([]int, l)

	L[0] = 1

	for i := 1; i < l; i++ {
		L[i] = nums[i-1] * L[i-1]
	}

	R[l-1] = 1
	for i := l - 2; i >= 0; i-- {
		R[i] = nums[i+1] * R[i+1]
	}

	for i := 0; i < l; i++ {
		ans[i] = L[i] * R[i]
	}

	return ans
}
