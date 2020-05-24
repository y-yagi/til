package main

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var count int
	dict := map[int]int{}

	for i := 0; i < len(A); i++ {
		for j := 0; j < len(B); j++ {
			v := A[i] + B[j]
			dict[v]++
		}
	}

	for i := 0; i < len(C); i++ {
		for j := 0; j < len(D); j++ {
			v := C[i] + D[j]
			if v, found := dict[-v]; found {
				count += v
			}
		}
	}

	return count
}
