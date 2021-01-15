package main

func wordBreak(s string, wordDict []string) []string {
	dict := make(map[string]bool)
	for _, w := range wordDict {
		dict[w] = true
	}

	mem := make(map[int][]string)
	return dfs(s, dict, mem)
}

func dfs(s string, dict map[string]bool, mem map[int][]string) []string {
	if v, exist := mem[len(s)]; exist {
		return v
	}

	retSlice := make([]string, 0)
	for i := 1; i <= len(s); i++ {
		if dict[s[:i]] {
			sliceV := dfs(s[i:], dict, mem)
			for _, v := range sliceV {
				retSlice = append(retSlice, s[:i]+" "+v)
			}
			if i == len(s) {
				retSlice = append(retSlice, s)
			}
		}
	}
	mem[len(s)] = retSlice
	return retSlice
}
