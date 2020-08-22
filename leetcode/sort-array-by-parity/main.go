package main

func sortArrayByParity(A []int) []int {
	i := 0
	j := len(A) - 1

	for i < j {
		for i < len(A)-1 && A[i]%2 == 0 {
			i++
		}

		for j >= 0 && A[j]%2 != 0 {
			j--
		}

		if i >= j {
			break
		}

		t := A[i]
		A[i] = A[j]
		A[j] = t

		i++
		j--
	}

	return A
}
