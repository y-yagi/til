package main

func validUtf8(data []int) bool {
	currentBits := 0
	var startBit uint8 = 7
	for _, v := range data {
		if currentBits == 0 {
			mask := (1 << startBit)
			for v&mask > 0 {
				startBit--
				mask = (1 << startBit)
				currentBits++
				if currentBits > 4 {
					return false
				}
			}

			if currentBits > 1 {
				currentBits--
			}
			startBit = 7
		} else if currentBits > 0 && v&(1<<7) > 0 && v&(1<<6) == 0 {
			currentBits--
		} else {
			return false
		}
	}
	if currentBits > 0 {
		return false
	}
	return true
}
