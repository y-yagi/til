package main

func shipWithinDays(weights []int, days int) int {
	left, right := 0, 0
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	for left < right {
		mid := (left + right) / 2
		if canShip(weights, days, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canShip(weights []int, days int, capacity int) bool {
	currWeight := 0
	daysNeeded := 1
	for _, w := range weights {
		currWeight += w
		if currWeight > capacity {
			daysNeeded++
			currWeight = w
		}
	}
	return daysNeeded <= days
}
