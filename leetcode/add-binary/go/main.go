package main

func addBinary(a string, b string) string {
	n := len(a)
	m := len(b)
	if n < m {
		return addBinary(b, a)
	}
	l := max(n, m)

	ans := []byte{}
	carry := 0
	j := m - 1

	for i := l - 1; i > -1; i-- {
		if a[i] == '1' {
			carry++
		}
		if j > -1 {
			if b[j] == '1' {
				carry++
			}
			j--
		}

		if carry%2 == 1 {
			ans = append([]byte{'1'}, ans...)
		} else {
			ans = append([]byte{'0'}, ans...)
		}
		carry /= 2
	}

	if carry == 1 {
		ans = append([]byte{'1'}, ans...)
	}

	return string(ans)
}

func max(i, j int) int {
	if i > j {
		return i
	}

	return j
}
