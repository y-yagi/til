package main

func subarraysDivByK(A []int, K int) int {
	dict := map[int]int{}
	dict[0] = 1
	var count, value int
	for _, a := range A {
		value = (value + a) % K
		if value < 0 {
			// To keep value as positive.
			value += K
		}
		count += dict[value]
		dict[value] = dict[value] + 1
	}
	return count
}
