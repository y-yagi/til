package main

func removeDuplicates(s string, k int) string {
	b := []byte(s)
	count := make([]int, len(b))
	for i := 0; i < len(b); i++ {
		if i == 0 || b[i] != b[i-1] {
			count[i] = 1
		} else {
			count[i] = count[i-1] + 1
			if count[i] == k {
				b = append(b[:i-k+1], b[i+1:]...)
				i = i - k
			}
		}
	}

	return string(b)
}
