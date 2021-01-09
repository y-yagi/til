package main

func numPairsDivisibleBy60(time []int) int {
	ans := 0
	count := map[int]int{}
	for _, t := range time {
		// (time[i] + time[j]) % 60 == (60 - t % 60) % 60
		ans += count[(60-t%60)%60]
		count[t%60]++
	}

	return ans
}
