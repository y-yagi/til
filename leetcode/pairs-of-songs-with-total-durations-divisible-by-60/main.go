package main

func numPairsDivisibleBy60(time []int) int {
	ans := 0
	remainders := [60]int{}
	for _, t := range time {
		if t%60 == 0 { // check if a%60==0 && b%60==0
			ans += remainders[0]
		} else { // check if a % 60+b%60==60
			ans += remainders[60-t%60]
		}
		remainders[t%60]++
	}
	return ans
}
