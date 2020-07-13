package main

func numJewelsInStones(J string, S string) int {
	stones := map[rune]int{}

	for _, v := range S {
		stones[v]++
	}

	ans := 0
	for _, v := range J {
		ans += stones[v]
	}
	return ans
}
