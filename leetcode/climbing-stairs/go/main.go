package main

func climbStairs(n int) int {
	if n == 1 {
		return n
	}

	t1 := 1
	t2 := 2
	t3 := 0

	for i := 3; i <= n; i++ {
		t3 = t2
		t2 = t1 + t2
		t1 = t3
	}

	return t2
}
