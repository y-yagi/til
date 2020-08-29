package main

func rand10() int {
	// NOTE: Using Rejection Sampling
	row := rand7()
	col := rand7()
	idx := col + (row-1)*7

	for idx > 40 {
		row = rand7()
		col = rand7()
		idx = col + (row-1)*7
	}

	return 1 + (idx-1)%10
}
