package main

type ArrayReader struct {
}

func (this *ArrayReader) get(index int) int { return 0 }

func search(reader ArrayReader, target int) int {
	top := 1
	for reader.get(top) < target {
		top = top << 1
	}

	bottom := top >> 1

	for bottom <= top {
		middle := (top + bottom) >> 1

		val := reader.get(middle)
		if val < target {
			bottom = middle + 1
		} else if val > target {
			top = middle - 1
		} else {
			return middle
		}
	}

	return -1
}
