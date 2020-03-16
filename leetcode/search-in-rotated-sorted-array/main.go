package main

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	lo := 0
	hi := len(nums) - 1

	for lo < hi {
		mid := (lo + hi) / 2
		if nums[mid] > nums[hi] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}

	rot := lo
	lo = 0
	hi = len(nums) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		realmid := (mid + rot) % len(nums)
		if nums[realmid] == target {
			return realmid
		}
		if nums[realmid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return -1
}
