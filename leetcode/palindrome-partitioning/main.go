package main

func partition(s string) [][]string {
	if len(s) == 0 {
		return nil
	}

	var res [][]string
	backtracking(&res, []string{}, s)
	return res
}

func backtracking(res *[][]string, prev []string, remain string) {
	if len(remain) == 0 {
		*res = append(*res, append([]string{}, prev...))
		return
	}

	for i := 0; i < len(remain); i++ {
		if !isPalindromeString(remain[:i+1]) {
			continue
		}
		prev = append(prev, remain[:i+1])
		backtracking(res, prev, remain[i+1:])
		prev = prev[:len(prev)-1]
	}
}

func isPalindromeString(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
