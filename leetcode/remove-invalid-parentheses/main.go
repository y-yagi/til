package main

func removeInvalidParentheses(s string) []string {
	l := 0
	r := 0
	res := []string{}
	for _, v := range s {
		if v == '(' {
			l++
		} else if v == ')' {
			if l == 0 {
				r++
			} else {
				l--
			}
		}
	}
	dfs(s, 0, l, r, &res)
	return res
}

func dfs(s string, start, l, r int, res *[]string) {
	if l == 0 && r == 0 {
		if isValid(s) {
			*res = append(*res, s)
		}
		return
	}
	for i := start; i < len(s); i++ {
		if i != start && s[i] == s[i-1] {
			continue
		}
		if s[i] == ')' && r > 0 {
			curr := s
			curr = curr[:i] + curr[i+1:]
			dfs(curr, i, l, r-1, res)
		} else if s[i] == '(' && l > 0 {
			curr := s
			curr = curr[:i] + curr[i+1:]
			dfs(curr, i, l-1, r, res)
		}
	}
}

func isValid(s string) bool {
	count := 0
	for _, v := range s {
		if v == '(' {
			count++
		}
		if v == ')' {
			count--
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}
