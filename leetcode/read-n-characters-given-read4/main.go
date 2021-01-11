package main

var solution = func(read4 func([]byte) int) func([]byte, int) int {
	return func(buf []byte, n int) int {
		b := make([]byte, 4)
		cnt, wrote := 0, 0
		for {
			if n-wrote >= 4 {
				cnt = read4(buf[wrote:])
				wrote += cnt
			} else {
				cnt = read4(b)
				for i := 0; i < cnt && wrote < n; i++ {
					buf[wrote] = b[i]
					wrote++
				}
			}
			if cnt == 0 {
				break
			}
		}
		return wrote
	}
}
