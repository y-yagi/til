package main

func duplicateZeros(arr []int) {
	l := len(arr) - 1
	dups := 0

	for i := 0; i <= l-dups; i++ {
		if arr[i] == 0 {
			if i == l-dups {
				arr[l] = 0
				l -= 1
				break
			}
			dups++
		}
	}

	last := l - dups
	for i := last; i >= 0; i-- {
		if arr[i] == 0 {
			arr[i+dups] = 0
			dups--
			arr[i+dups] = 0
		} else {
			arr[i+dups] = arr[i]
		}
	}
}
