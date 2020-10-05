package main

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}

	// This comes from `highestOneBit`.
	bitmask := N
	bitmask |= (bitmask >> 1)
	bitmask |= (bitmask >> 2)
	bitmask |= (bitmask >> 4)
	bitmask |= (bitmask >> 8)
	bitmask |= (bitmask >> 16)
	// flip all bits
	return bitmask ^ N
}
