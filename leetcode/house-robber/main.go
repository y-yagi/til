package main

func rob(nums []int) int {
	prevMax, currMax := 0, 0
	for _, num := range nums {
		t := currMax
		// f(k) = max(f(k – 2) + Ak, f(k – 1))
		currMax = max(prevMax+num, currMax)
		prevMax = t
	}

	return currMax
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
