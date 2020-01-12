package main

func convertToTitle(n int) string {
	a := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	l := len(a)
	answer := ""

	for n > l {
		m := n % l
		n = n / l

		if m == 0 {
			answer = string(a[l-1]) + answer
			n--
		} else {
			answer = string(a[m-1]) + answer
		}
	}
	answer = string(a[n-1]) + answer

	return answer
}
