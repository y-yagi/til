package main

func multiply(num1 string, num2 string) string {
	l1 := len(num1)
	l2 := len(num2)
	pos := make([]int, l1+l2)

	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			mul := int(num1[i]-'0') * int(num2[j]-'0')
			p1 := i + j
			p2 := i + j + 1
			sum := mul + pos[p2]

			pos[p1] += sum / 10
			pos[p2] = (sum) % 10
		}
	}

	sb := make([]byte, 0)
	for _, p := range pos {
		if !(len(sb) == 0 && p == 0) {
			sb = append(sb, byte(p+'0'))
		}
	}

	if len(sb) == 0 {
		return "0"
	}
	return string(sb)
}
