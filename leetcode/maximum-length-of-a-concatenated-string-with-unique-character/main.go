package main

func maxLength(arr []string) int {
	if len(arr) == 1 {
		return len(arr[0])
	}

	max := 0
	dfs(arr, 0, "", &max)
	return max
}

func dfs(arr []string, index int, cur string, max *int) {
	if index == len(arr) && uniqueCharCount(cur) > *max {
		*max = len(cur)
		return
	}

	if index == len(arr) {
		return
	}

	if uniqueCharCount(cur) == -1 {
		return
	}

	dfs(arr, index+1, cur+arr[index], max)
	dfs(arr, index+1, cur, max)
}

func uniqueCharCount(s string) int {
	chars := make([]int, 26)
	for _, c := range s {
		if chars[c-'a'] == 1 {
			return -1
		}
		chars[c-'a']++
	}
	return len(s)
}
