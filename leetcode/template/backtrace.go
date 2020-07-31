package main

func bactrace(candidate int) {
	if findSolution(candidate) {
		output(candidate)
	}

	for _, next_candidate := range list_of_candidate {
		if isValid(next_candidate) {
			// try this partial candidate solution
			place(next_candidate)
			// given the candidate, explore further.
			backtrack(next_candidate)
			// backtrack
			remove(next_candidate)
		}
	}
}
