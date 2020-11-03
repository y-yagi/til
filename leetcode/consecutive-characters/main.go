package main

func maxPower(s string) int {
	last := s[0]
	count := 1
	ans := 1

	for i := 1; i < len(s); i++ {
		if s[i] == last {
			count++
			continue
		}
		ans = max(ans, count)
		count = 1
		last = s[i]
	}

	return max(ans, count)
}

func max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
