package main

func generateParenthesis(n int) []string {
	return backtrace([]string{}, "", 0, 0, n)
}

func backtrace(patterns []string, cur string, open, close, max int) []string {
	if len(cur) == max*2 {
		patterns = append(patterns, cur)
		return patterns
	}

	if open < max {
		patterns = backtrace(patterns, cur+"(", open+1, close, max)
	}
	if close < open {
		patterns = backtrace(patterns, cur+")", open, close+1, max)
	}
	return patterns
}
