package main

func splitArray(nums []int, m int) int {
	l := 0
	r := 0
	n := len(nums)
	for i := 0; i < n; i++ {
		r += nums[i]
		if l < nums[i] {
			l = nums[i]
		}
	}

	ans := r
	for l <= r {
		mid := (l + r) / 2
		sum := 0
		cnt := 1
		for i := 0; i < n; i++ {
			if sum+nums[i] > mid {
				cnt++
				sum = nums[i]
			} else {
				sum += nums[i]
			}
		}

		if cnt <= m {
			ans = min(ans, mid)
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return ans
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
