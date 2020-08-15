package main

func longestPalindrome(s string) int {
	count := make([]int, 128)
	for _, b := range s {
		index := b - '0'
		count[index]++
	}

	ans := 0
	for _, v := range count {
		ans += v / 2 * 2
		if ans%2 == 0 && v%2 == 1 {
			ans++
		}
	}

	return ans
}
