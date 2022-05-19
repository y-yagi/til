package main

func hammingWeight(num uint32) int {
	var answer int

	for i := 0; i < 32; i++ {
		if (num & 1) == 1 {
			answer++
		}
		num >>= 1
	}

	return answer
}
