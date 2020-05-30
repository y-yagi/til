package main

func calculate(s string) int {
	stack := []int{}
	var num int
	var sign byte
	sign = '+'

	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			num = num*10 + int(s[i]-'0')
		}

		if s[i] == '*' || s[i] == '/' || s[i] == '+' || s[i] == '-' || i == len(s)-1 {
			if sign == '-' {
				stack = append(stack, -num)
			}
			if sign == '+' {
				stack = append(stack, num)
			}
			if sign == '*' {
				x := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				stack = append(stack, x*num)
			}
			if sign == '/' {
				x := stack[len(stack)-1]
				stack = stack[0 : len(stack)-1]
				stack = append(stack, x/num)
			}
			sign = s[i]
			num = 0
		}
	}

	ans := 0
	for _, v := range stack {
		ans += v
	}

	return ans
}
