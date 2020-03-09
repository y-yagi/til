package main

func checkPerfectNumber(num int) bool {
	sum := 1

	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			if i*i != num {
				sum = sum + i + num/i
			} else {
				sum = sum + i
			}
		}
	}

	if sum == num && sum != 1 {
		return true
	}
	return false
}
