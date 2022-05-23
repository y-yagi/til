package main

func backspaceCompare(S string, T string) bool {
	stackS := []byte{}
	stackT := []byte{}

	for i := 0; i < len(S); i++ {
		if S[i] == '#' {
			if len(stackS) > 0 {
				stackS = stackS[0 : len(stackS)-1]
			}
		} else {
			stackS = append(stackS, S[i])
		}
	}

	for i := 0; i < len(T); i++ {
		if T[i] == '#' {
			if len(stackT) > 0 {
				stackT = stackT[0 : len(stackT)-1]
			}
		} else {
			stackT = append(stackT, T[i])
		}
	}

	return string(stackS) == string(stackT)
}
