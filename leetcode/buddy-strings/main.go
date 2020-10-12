package main

func buddyStrings(A string, B string) bool {
	if len(A) != len(B) {
		return false
	}

	if A == B {
		count := make([]int, 26)
		for i := 0; i < len(A); i++ {
			count[A[i]-'a']++
		}
		for _, c := range count {
			if c > 1 {
				return true
			}
		}
		return false
	}

	first := -1
	second := -1
	for i := 0; i < len(A); i++ {
		if A[i] != B[i] {
			if first == -1 {
				first = i
			} else if second == -1 {
				second = i
			} else {
				return false
			}
		}
	}

	return second != -1 && A[first] == B[second] && A[second] == B[first]
}
