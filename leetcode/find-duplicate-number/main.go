package main

func findDuplicate(nums []int) int {
	// Find the intersection point of the two runners
	tortoise := nums[0]
	hare := nums[0]

	for {
		tortoise = nums[tortoise]
		hare = nums[nums[hare]]
		if tortoise == hare {
			break
		}
	}

	// Find the "entrance" the cycle
	tortoise = nums[0]
	for tortoise != hare {
		tortoise = nums[tortoise]
		hare = nums[hare]
	}

	return hare
}
