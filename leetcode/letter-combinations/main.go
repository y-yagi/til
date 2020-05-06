package main

var dict = map[string]string{"2": "abc", "3": "def", "4": "ghi", "5": "jkl", "6": "mno", "7": "pqrs", "8": "tuv", "9": "wxyz"}

func letterCombinations(digits string) []string {
	var ans []string
	if len(digits) != 0 {
		ans = backtrace("", digits, ans)
	}

	return ans
}

func backtrace(combination, nextDigit string, ans []string) []string {
	if len(nextDigit) == 0 {
		ans = append(ans, combination)
	} else {
		digit := nextDigit[0:1]
		lettters := dict[digit]

		for i := 0; i < len(lettters); i++ {
			letter := dict[digit][i : i+1]
			ans = backtrace(combination+letter, nextDigit[1:], ans)
		}
	}
	return ans
}
