package main

func kthGrammar(N int, K int) int {
	if N == 1 {
		return 0
	}
	if K%2 == 0 {
		if kthGrammar(N-1, K/2) == 0 {
			return 1
		} else {
			return 0
		}
	} else {
		if kthGrammar(N-1, (K+1)/2) == 0 {
			return 0
		} else {
			return 1
		}
	}
}
