package main

func shipWithinDays(weights []int, D int) int {
	var left, right int
	for _, w := range weights {
		left = max(left, w)
		right += w
	}

	for left < right {
		mid := (left + right) / 2
		need := 1
		cur := 0

		for _, w := range weights {
			if cur+w > mid {
				need += 1
				cur = 0
			}
			cur += w
		}

		if need > D {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return left
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
