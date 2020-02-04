package main

func isUgly(num int) bool {
	if num < 1 {
		return false
	}
	factors := [3]int{2, 3, 5}
	for _, v := range factors {
		for num%v == 0 {
			num /= v
		}
	}
	return num == 1
}
