package main

func findPoisonedDuration(timeSeries []int, duration int) int {
	if len(timeSeries) == 0 {
		return 0
	}

	total := 0
	for i := 0; i < len(timeSeries)-1; i++ {
		total += min(timeSeries[i+1]-timeSeries[i], duration)
	}

	return total + duration
}

func min(x, y int) int {
	if x > y {
		return y
	}

	return x
}
