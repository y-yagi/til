package main

func numWays(n int, k int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return k
	}

	// First 2 posts have same color.
	sameCase := k
	// First 2 posts have different colors.
	diffCase := k * (k - 1)

	for i := 2; i < n; i++ {
		t := diffCase
		// To every same case and diff case we can add a new post with different color as the last one.
		// We have k-1 color options for the last one.
		diffCase = (diffCase + sameCase) * (k - 1)
		// To every diff case can add a new post with the same color as the last one to not generate violation,
		// no more than 2 adjacent fence posts have the same color.
		sameCase = t
	}

	return diffCase + sameCase
}
