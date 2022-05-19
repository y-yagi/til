package main

func reverseBits(num uint32) uint32 {
	var answer uint32

	for i := 0; i < 32; i++ {
		answer <<= 1
		if (num & 1) == 1 {
			answer++
		}
		num >>= 1
	}

	return answer
}
