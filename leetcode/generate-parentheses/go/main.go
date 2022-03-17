package main

func generateParenthesis(n int) []string {
	ans := []string{}
	backtrace(&ans, "", 0, 0, n)
	return ans
}

func backtrace(ans *[]string, cur string, opened, closed, max int) {
	if len(cur) == max*2 {
		*ans = append(*ans, cur)
		return
	}

	if opened < max {
		backtrace(ans, cur+"(", opened+1, closed, max)
	}
	if closed < opened {
		backtrace(ans, cur+")", opened, closed+1, max)
	}
}
