package main

func isHappy(n int) bool {
	stack := []int{n}
	return judge(n, stack)
}

func judge(n int, stack []int) bool {
	v := 0

	for n > 0 {
		v += (n % 10) * (n % 10)
		n = n / 10
	}

	if v == 1 {
		return true
	}

	if contains(stack, v) {
		return false
	}

	stack = append(stack, v)
	return judge(v, stack)
}

func contains(s []int, e int) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}
